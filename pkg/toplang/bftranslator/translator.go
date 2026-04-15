package bftranslator

import (
	"brainfuck-compiler-go/pkg/toplang/parser"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const defaultMaxCells = 1024

// Translator converts TopLang source code into a brainfuck program string.
type Translator struct {
	mem *memory
	cg  *codegen
	err error
}

// NewTranslator creates a new Translator.
func NewTranslator() *Translator {
	return &Translator{}
}

// Translate parses TopLang source and returns the equivalent brainfuck program.
func (t *Translator) Translate(source string) (string, error) {
	t.mem = newMemory(defaultMaxCells)
	t.cg = &codegen{}
	t.err = nil

	tree, _, _, err := t.parse(source)
	if err != nil {
		return "", err
	}

	t.visitProgram(tree.(*parser.ProgramContext))
	if t.err != nil {
		return "", t.err
	}

	return t.cg.String(), nil
}

// GetParseTree returns the ANTLR parse tree for AST visualisation.
func (t *Translator) GetParseTree(source string) (antlr.Tree, []string, error) {
	tree, _, p, err := t.parse(source)
	if err != nil {
		return nil, nil, err
	}

	return tree, p.GetRuleNames(), nil
}

// VisualizeAST returns a human-readable indented tree of the ANTLR parse tree.
func (t *Translator) VisualizeAST(source string) (string, error) {
	tree, ruleNames, err := t.GetParseTree(source)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	visualizeTree(&sb, tree, ruleNames, "", true)

	return sb.String(), nil
}

func visualizeTree(sb *strings.Builder, tree antlr.Tree, ruleNames []string, prefix string, isLast bool) {
	connector := "├── "
	childPrefix := "│   "
	if isLast {
		connector = "└── "
		childPrefix = "    "
	}

	if prefix == "" && isLast {
		connector = ""
		childPrefix = ""
	}

	sb.WriteString(prefix)
	sb.WriteString(connector)

	switch node := tree.(type) {
	case antlr.ParserRuleContext:
		ruleIdx := node.GetRuleIndex()
		name := fmt.Sprintf("%s", ruleNames[ruleIdx])
		sb.WriteString(name)
	case antlr.TerminalNode:
		sb.WriteString(fmt.Sprintf("'%s'", node.GetText()))
	default:
		sb.WriteString(fmt.Sprintf("%v", tree))
	}
	sb.WriteByte('\n')

	children := tree.GetChildren()
	for i, child := range children {
		visualizeTree(sb, child.(antlr.Tree), ruleNames, prefix+childPrefix, i == len(children)-1)
	}
}

func (t *Translator) parse(source string) (antlr.ParserRuleContext, *antlr.CommonTokenStream, *parser.TopLangParser, error) {
	input := antlr.NewInputStream(source)
	lexer := parser.NewTopLangLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTopLangParser(tokens)

	el := &errCollector{}
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	tree := p.Program()
	if len(el.errs) > 0 {
		return nil, nil, nil, fmt.Errorf("%w: %s", ErrParseFailed, el.errs[0])
	}

	return tree.(*parser.ProgramContext), tokens, p, nil
}

func (t *Translator) visitProgram(ctx *parser.ProgramContext) {
	for _, s := range ctx.AllStatement() {
		if t.err != nil {
			return
		}
		t.visitStatement(s.(*parser.StatementContext))
	}
}

func (t *Translator) visitStatement(ctx *parser.StatementContext) {
	if t.err != nil {
		return
	}

	switch {
	case ctx.VarDecl() != nil:
		t.visitVarDecl(ctx.VarDecl().(*parser.VarDeclContext))
	case ctx.Assignment() != nil:
		t.visitAssignment(ctx.Assignment().(*parser.AssignmentContext))
	case ctx.PrintStmt() != nil:
		t.visitPrintStmt(ctx.PrintStmt().(*parser.PrintStmtContext))
	case ctx.PrintNumberStmt() != nil:
		t.visitPrintNumberStmt(ctx.PrintNumberStmt().(*parser.PrintNumberStmtContext))
	case ctx.PrintStringStmt() != nil:
		t.visitPrintStringStmt(ctx.PrintStringStmt().(*parser.PrintStringStmtContext))
	case ctx.ReadStmt() != nil:
		t.visitReadStmt(ctx.ReadStmt().(*parser.ReadStmtContext))
	case ctx.IfStmt() != nil:
		t.visitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	case ctx.WhileStmt() != nil:
		t.visitWhileStmt(ctx.WhileStmt().(*parser.WhileStmtContext))
	case ctx.Block() != nil:
		t.visitBlock(ctx.Block().(*parser.BlockContext))
	}
}

func (t *Translator) visitVarDecl(ctx *parser.VarDeclContext) {
	name := ctx.ID().GetText()

	cell, err := t.mem.declareVar(name)
	if err != nil {
		t.err = fmt.Errorf("%w: %s", err, name)
		return
	}

	val := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	t.cg.clearCell(cell)
	t.cg.moveValue(val, cell)
	t.mem.freeTemp(val)
}

func (t *Translator) visitAssignment(ctx *parser.AssignmentContext) {
	name := ctx.ID().GetText()

	cell, err := t.mem.lookupVar(name)
	if err != nil {
		t.err = fmt.Errorf("%w: %s", err, name)
		return
	}

	val := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	t.cg.clearCell(cell)
	t.cg.moveValue(val, cell)
	t.mem.freeTemp(val)
}

func (t *Translator) visitPrintStmt(ctx *parser.PrintStmtContext) {
	val := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	t.cg.moveTo(val)
	t.cg.emit(".")
	t.cg.clearCell(val)
	t.mem.freeTemp(val)
}

// visitPrintNumberStmt prints the byte value (0-255) as its decimal digits.
// Splits value into hundreds/tens/ones via two divmods, suppresses leading zeros,
// adds ASCII '0' offset (48) to each digit before emitting '.'.
func (t *Translator) visitPrintNumberStmt(ctx *parser.PrintNumberStmtContext) {
	val := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	const100 := t.allocOrFail()
	const10 := t.allocOrFail()
	if t.err != nil {
		return
	}
	t.cg.setCell(const100, 100)
	t.cg.setCell(const10, 10)

	hundreds, tensMod := t.emitDivMod(val, const100)
	if t.err != nil {
		return
	}

	tens, ones := t.emitDivMod(tensMod, const10)
	if t.err != nil {
		return
	}

	nonLeading := t.allocOrFail()
	scratch := t.allocOrFail()
	if t.err != nil {
		return
	}
	t.cg.clearCell(nonLeading)
	t.cg.clearCell(scratch)

	// if hundreds > 0 { print digit; nonLeading = 1 }
	t.cg.moveTo(hundreds)
	t.cg.emit("[")
	t.cg.emit(strings.Repeat("+", 48))
	t.cg.emit(".")
	t.cg.clearCell(hundreds)
	t.cg.moveTo(nonLeading)
	t.cg.emit("+")
	t.cg.moveTo(hundreds)
	t.cg.emit("]")

	// nonLeading += tens (preserves tens); now nonLeading != 0 iff we should print tens digit
	t.cg.clearCell(scratch)
	t.cg.copyCell(tens, nonLeading, scratch)

	// if nonLeading > 0 { print tens digit }
	t.cg.moveTo(nonLeading)
	t.cg.emit("[")
	t.cg.moveTo(tens)
	t.cg.emit(strings.Repeat("+", 48))
	t.cg.emit(".")
	t.cg.clearCell(tens)
	t.cg.clearCell(nonLeading)
	t.cg.emit("]")

	// always print ones digit
	t.cg.moveTo(ones)
	t.cg.emit(strings.Repeat("+", 48))
	t.cg.emit(".")
	t.cg.clearCell(ones)

	t.mem.freeTemp(hundreds)
	t.mem.freeTemp(tens)
	t.mem.freeTemp(ones)
	t.mem.freeTemp(nonLeading)
	t.mem.freeTemp(scratch)
}

// visitPrintStringStmt emits BF to print a string literal.
// Uses a single scratch cell and emits the minimal delta (+/-) between
// successive character values before each '.', wrapping mod 256 to pick
// the shorter direction.
func (t *Translator) visitPrintStringStmt(ctx *parser.PrintStringStmtContext) {
	raw := ctx.STRING().GetText()
	s, err := strconv.Unquote(raw)
	if err != nil {
		t.err = fmt.Errorf("%w: %s", ErrInvalidString, raw)
		return
	}

	if len(s) == 0 {
		return
	}

	scratch := t.allocOrFail()
	if t.err != nil {
		return
	}
	t.cg.clearCell(scratch)

	prev := 0
	for _, ch := range []byte(s) {
		curr := int(ch)
		forward := (curr - prev + 256) % 256
		if forward <= 128 {
			t.cg.emit(strings.Repeat("+", forward))
		} else {
			t.cg.emit(strings.Repeat("-", 256-forward))
		}
		t.cg.emit(".")
		prev = curr
	}

	t.cg.clearCell(scratch)
	t.mem.freeTemp(scratch)
}

func (t *Translator) visitReadStmt(ctx *parser.ReadStmtContext) {
	name := ctx.ID().GetText()

	cell, err := t.mem.lookupVar(name)
	if err != nil {
		t.err = fmt.Errorf("%w: %s", err, name)
		return
	}

	t.cg.clearCell(cell)
	t.cg.moveTo(cell)
	t.cg.emit(",")
}

func (t *Translator) visitIfStmt(ctx *parser.IfStmtContext) {
	cond := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	blocks := ctx.AllBlock()
	hasElse := len(blocks) > 1

	if hasElse {
		tmp := t.allocOrFail()
		if t.err != nil {
			return
		}

		t.cg.clearCell(tmp)
		t.cg.emitIfElse(cond, tmp,
			func() {
				t.visitBlock(blocks[0].(*parser.BlockContext))
			},
			func() {
				t.visitBlock(blocks[1].(*parser.BlockContext))
			},
		)
		t.mem.freeTemp(tmp)
	} else {
		t.cg.moveTo(cond)
		t.cg.emit("[")
		t.visitBlock(blocks[0].(*parser.BlockContext))
		t.cg.clearCell(cond)
		t.cg.emit("]")
	}
	t.mem.freeTemp(cond)
}

func (t *Translator) visitWhileStmt(ctx *parser.WhileStmtContext) {
	cond := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}

	t.cg.moveTo(cond)
	t.cg.emit("[")

	t.visitBlock(ctx.Block().(*parser.BlockContext))

	// re-evaluate condition into the SAME cell SO THAT WE DONT USE AN ADDITIONAL ONE!!
	t.cg.clearCell(cond)
	t.mem.freeTemp(cond)

	newCond := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return
	}
	if newCond != cond {
		t.cg.clearCell(cond)
		t.cg.moveValue(newCond, cond)
	}
	t.mem.freeTemp(newCond)

	t.cg.moveTo(cond)
	t.cg.emit("]")
}

func (t *Translator) visitBlock(ctx *parser.BlockContext) {
	for _, s := range ctx.AllStatement() {
		if t.err != nil {
			return
		}

		t.visitStatement(s.(*parser.StatementContext))
	}
}

// visitor expressions
// Each visitExpr returns a temp cell holding the result.
// Caller is responsible for freeing it.

func (t *Translator) visitExpr(ctx parser.IExprContext) int {
	if t.err != nil {
		return 0
	}

	switch c := ctx.(type) {
	case *parser.IntLiteralContext:
		return t.visitIntLiteral(c)
	case *parser.IdExprContext:
		return t.visitIdExpr(c)
	case *parser.ParenExprContext:
		return t.visitExpr(c.Expr())
	case *parser.UnaryMinusExprContext:
		return t.visitUnaryMinus(c)
	case *parser.AddExprContext:
		return t.visitAddExpr(c)
	case *parser.MulExprContext:
		return t.visitMulExpr(c)
	case *parser.CompExprContext:
		return t.visitCompExpr(c)
	default:
		t.err = fmt.Errorf("unknown expression type: %T", ctx)
		return 0
	}
}

func (t *Translator) visitIntLiteral(ctx *parser.IntLiteralContext) int {
	v, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		t.err = fmt.Errorf("%w: %s", ErrIntOverflow, ctx.INT().GetText())
		return 0
	}

	v = v % 256

	cell := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.setCell(cell, v)
	return cell
}

func (t *Translator) visitIdExpr(ctx *parser.IdExprContext) int {
	name := ctx.ID().GetText()

	src, err := t.mem.lookupVar(name)
	if err != nil {
		t.err = fmt.Errorf("%w: %s", err, name)
		return 0
	}

	dst := t.allocOrFail()
	tmp := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(dst)
	t.cg.clearCell(tmp)
	t.cg.copyCell(src, dst, tmp)
	t.mem.freeTemp(tmp)

	return dst
}

func (t *Translator) visitUnaryMinus(ctx *parser.UnaryMinusExprContext) int {
	inner := t.visitExpr(ctx.Expr())
	if t.err != nil {
		return 0
	}

	result := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(result)
	// result = 0 - inner // inner[ result- inner- ] // loop and subtract inner while inner > 0
	t.cg.moveTo(inner)
	t.cg.emit("[")
	t.cg.moveTo(result)
	t.cg.emit("-")
	t.cg.moveTo(inner)
	t.cg.emit("-]")
	t.mem.freeTemp(inner)
	return result
}

func (t *Translator) visitAddExpr(ctx *parser.AddExprContext) int {
	l := t.visitExpr(ctx.GetLeft())
	if t.err != nil {
		return 0
	}

	r := t.visitExpr(ctx.GetRight())
	if t.err != nil {
		return 0
	}

	res := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(res)
	switch ctx.GetOp().GetText() {
	case "+":
		t.cg.addCells(l, r, res)
	case "-":
		t.cg.subCells(l, r, res)
	default:
		t.err = fmt.Errorf("%w: %s", ErrUnsupportedOperator, ctx.GetOp().GetText())
		return 0
	}

	t.mem.freeTemp(l)
	t.mem.freeTemp(r)
	return res
}

func (t *Translator) visitMulExpr(ctx *parser.MulExprContext) int {
	l := t.visitExpr(ctx.GetLeft())
	if t.err != nil {
		return 0
	}

	r := t.visitExpr(ctx.GetRight())
	if t.err != nil {
		return 0
	}

	op := ctx.GetOp().GetText()
	switch op {
	case "*":
		return t.emitMul(l, r)
	case "/":
		return t.emitDiv(l, r)
	case "%":
		return t.emitMod(l, r)
	default:
		t.err = fmt.Errorf("%w: %s", ErrUnsupportedOperator, op)
		return 0
	}
}

func (t *Translator) emitMul(a, b int) int {
	res := t.allocOrFail()
	tmp := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(res)
	t.cg.clearCell(tmp)
	t.cg.multiplyCells(a, b, res, tmp)
	t.mem.freeTemp(a)
	t.mem.freeTemp(b)
	t.mem.freeTemp(tmp)

	return res
}

func (t *Translator) emitDiv(a, b int) int {
	q, r := t.emitDivMod(a, b)
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(r)
	t.mem.freeTemp(r)

	return q
}

func (t *Translator) emitMod(a, b int) int {
	q, r := t.emitDivMod(a, b)
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(q)
	t.mem.freeTemp(q)

	return r
}

func (t *Translator) emitDivMod(a, b int) (quotient, remainder int) {
	quotient = t.allocOrFail()
	remainder = t.allocOrFail()
	rcopy := t.allocOrFail()
	counter := t.allocOrFail()
	success := t.allocOrFail()
	s1 := t.allocOrFail()
	s2 := t.allocOrFail()
	s3 := t.allocOrFail()
	if t.err != nil {
		return
	}

	t.cg.clearCell(quotient)
	t.cg.clearCell(remainder)
	t.cg.clearCell(rcopy)
	t.cg.clearCell(counter)
	t.cg.clearCell(success)
	t.cg.clearCell(s1)
	t.cg.clearCell(s2)
	t.cg.clearCell(s3)

	t.cg.emitDivMod(a, b, quotient, remainder, rcopy, counter, success, s1, s2, s3)

	t.mem.freeTemp(a)
	t.mem.freeTemp(b)
	t.mem.freeTemp(rcopy)
	t.mem.freeTemp(counter)
	t.mem.freeTemp(success)
	t.mem.freeTemp(s1)
	t.mem.freeTemp(s2)
	t.mem.freeTemp(s3)
	return
}

// visitor: comparisons

func (t *Translator) visitCompExpr(ctx *parser.CompExprContext) int {
	l := t.visitExpr(ctx.GetLeft())
	if t.err != nil {
		return 0
	}

	r := t.visitExpr(ctx.GetRight())
	if t.err != nil {
		return 0
	}

	switch ctx.GetOp().GetText() {
	case "==":
		return t.emitEq(l, r)
	case "!=":
		return t.emitNeq(l, r)
	case "<":
		return t.emitLT(l, r)
	case ">":
		return t.emitLT(r, l) // a > b equals b < a
	case "<=":
		return t.emitLTE(l, r)
	case ">=":
		return t.emitLTE(r, l) // a >= b equals b <= a
	default:
		t.err = fmt.Errorf("%w: %s", ErrUnsupportedOperator, ctx.GetOp().GetText())
		return 0
	}
}

func (t *Translator) emitEq(a, b int) int {
	diff := t.allocOrFail()
	res := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(diff)
	t.cg.clearCell(res)
	t.cg.subCells(a, b, diff)
	t.mem.freeTemp(a)
	t.mem.freeTemp(b)
	// result = 1 if diff==0
	t.cg.setCell(res, 1)
	t.cg.moveTo(diff)
	t.cg.emit("[")
	t.cg.clearCell(res)
	t.cg.clearCell(diff)
	t.cg.emit("]")
	t.mem.freeTemp(diff)

	return res
}

func (t *Translator) emitNeq(a, b int) int {
	diff := t.allocOrFail()
	res := t.allocOrFail()
	tmp := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(diff)
	t.cg.clearCell(res)
	t.cg.clearCell(tmp)
	t.cg.subCells(a, b, diff)
	t.mem.freeTemp(a)
	t.mem.freeTemp(b)
	// res = bool(diff)
	t.cg.normalizeBool(diff, tmp)
	t.cg.moveValue(diff, res)
	t.mem.freeTemp(diff)
	t.mem.freeTemp(tmp)

	return res
}

func (t *Translator) emitLT(a, b int) int {
	res := t.allocOrFail()
	s1 := t.allocOrFail()
	s2 := t.allocOrFail()
	s3 := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	t.cg.clearCell(res)
	t.cg.clearCell(s1)
	t.cg.clearCell(s2)
	t.cg.clearCell(s3)
	t.cg.emitLessThan(a, b, res, s1, s2, s3)
	t.mem.freeTemp(a)
	t.mem.freeTemp(b)
	t.mem.freeTemp(s1)
	t.mem.freeTemp(s2)
	t.mem.freeTemp(s3)
	return res
}

func (t *Translator) emitLTE(a, b int) int {
	// a <= b equals !(b < a)
	lt := t.emitLT(b, a)
	if t.err != nil {
		return 0
	}

	res := t.allocOrFail()
	if t.err != nil {
		return 0
	}

	// negate: res = (lt == 0) ? 1 : 0
	t.cg.clearCell(res)
	t.cg.setCell(res, 1)
	t.cg.moveTo(lt)
	t.cg.emit("[")
	t.cg.clearCell(res)
	t.cg.clearCell(lt)
	t.cg.emit("]")
	t.mem.freeTemp(lt)
	return res
}

// helper

func (t *Translator) allocOrFail() int {
	if t.err != nil {
		return 0
	}

	c, err := t.mem.allocTemp()
	if err != nil {
		t.err = err
		return 0
	}

	return c
}
