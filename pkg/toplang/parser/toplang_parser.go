// Code generated from tools/TopLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TopLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TopLangParser struct {
	*antlr.BaseParser
}

var TopLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func toplangParserInit() {
	staticData := &TopLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'var'", "'print'", "'printNumber'", "'printString'", "'read'",
		"'if'", "'else'", "'while'", "'=='", "'!='", "'<='", "'>='", "'<'",
		"'>'", "'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'('", "')'", "'{'",
		"'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "VAR", "PRINT", "PRINTNUMBER", "PRINTSTRING", "READ", "IF", "ELSE",
		"WHILE", "EQ", "NEQ", "LTE", "GTE", "LT", "GT", "PLUS", "MINUS", "MUL",
		"DIV", "MOD", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "SEMI",
		"INT", "STRING", "ID", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "varDecl", "assignment", "printStmt", "printNumberStmt",
		"printStringStmt", "readStmt", "ifStmt", "whileStmt", "block", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 128, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 42, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 86, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 96, 8, 10, 10, 10, 12, 10, 99, 9, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 112, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 123, 8, 11, 10, 11, 12, 11, 126, 9, 11, 1, 11, 0,
		1, 22, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 3, 1, 0, 9, 14,
		1, 0, 15, 16, 1, 0, 17, 19, 132, 0, 27, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0,
		4, 43, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 60, 1, 0,
		0, 0, 12, 66, 1, 0, 0, 0, 14, 72, 1, 0, 0, 0, 16, 78, 1, 0, 0, 0, 18, 87,
		1, 0, 0, 0, 20, 93, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0,
		25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1,
		0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31,
		1, 1, 0, 0, 0, 32, 42, 3, 4, 2, 0, 33, 42, 3, 6, 3, 0, 34, 42, 3, 8, 4,
		0, 35, 42, 3, 10, 5, 0, 36, 42, 3, 12, 6, 0, 37, 42, 3, 14, 7, 0, 38, 42,
		3, 16, 8, 0, 39, 42, 3, 18, 9, 0, 40, 42, 3, 20, 10, 0, 41, 32, 1, 0, 0,
		0, 41, 33, 1, 0, 0, 0, 41, 34, 1, 0, 0, 0, 41, 35, 1, 0, 0, 0, 41, 36,
		1, 0, 0, 0, 41, 37, 1, 0, 0, 0, 41, 38, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		41, 40, 1, 0, 0, 0, 42, 3, 1, 0, 0, 0, 43, 44, 5, 1, 0, 0, 44, 45, 5, 28,
		0, 0, 45, 46, 5, 20, 0, 0, 46, 47, 3, 22, 11, 0, 47, 48, 5, 25, 0, 0, 48,
		5, 1, 0, 0, 0, 49, 50, 5, 28, 0, 0, 50, 51, 5, 20, 0, 0, 51, 52, 3, 22,
		11, 0, 52, 53, 5, 25, 0, 0, 53, 7, 1, 0, 0, 0, 54, 55, 5, 2, 0, 0, 55,
		56, 5, 21, 0, 0, 56, 57, 3, 22, 11, 0, 57, 58, 5, 22, 0, 0, 58, 59, 5,
		25, 0, 0, 59, 9, 1, 0, 0, 0, 60, 61, 5, 3, 0, 0, 61, 62, 5, 21, 0, 0, 62,
		63, 3, 22, 11, 0, 63, 64, 5, 22, 0, 0, 64, 65, 5, 25, 0, 0, 65, 11, 1,
		0, 0, 0, 66, 67, 5, 4, 0, 0, 67, 68, 5, 21, 0, 0, 68, 69, 5, 27, 0, 0,
		69, 70, 5, 22, 0, 0, 70, 71, 5, 25, 0, 0, 71, 13, 1, 0, 0, 0, 72, 73, 5,
		5, 0, 0, 73, 74, 5, 21, 0, 0, 74, 75, 5, 28, 0, 0, 75, 76, 5, 22, 0, 0,
		76, 77, 5, 25, 0, 0, 77, 15, 1, 0, 0, 0, 78, 79, 5, 6, 0, 0, 79, 80, 5,
		21, 0, 0, 80, 81, 3, 22, 11, 0, 81, 82, 5, 22, 0, 0, 82, 85, 3, 20, 10,
		0, 83, 84, 5, 7, 0, 0, 84, 86, 3, 20, 10, 0, 85, 83, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 17, 1, 0, 0, 0, 87, 88, 5, 8, 0, 0, 88, 89, 5, 21, 0, 0,
		89, 90, 3, 22, 11, 0, 90, 91, 5, 22, 0, 0, 91, 92, 3, 20, 10, 0, 92, 19,
		1, 0, 0, 0, 93, 97, 5, 23, 0, 0, 94, 96, 3, 2, 1, 0, 95, 94, 1, 0, 0, 0,
		96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1,
		0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 101, 5, 24, 0, 0, 101, 21, 1, 0, 0, 0,
		102, 103, 6, 11, -1, 0, 103, 104, 5, 16, 0, 0, 104, 112, 3, 22, 11, 4,
		105, 106, 5, 21, 0, 0, 106, 107, 3, 22, 11, 0, 107, 108, 5, 22, 0, 0, 108,
		112, 1, 0, 0, 0, 109, 112, 5, 26, 0, 0, 110, 112, 5, 28, 0, 0, 111, 102,
		1, 0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0,
		0, 0, 112, 124, 1, 0, 0, 0, 113, 114, 10, 7, 0, 0, 114, 115, 7, 0, 0, 0,
		115, 123, 3, 22, 11, 8, 116, 117, 10, 6, 0, 0, 117, 118, 7, 1, 0, 0, 118,
		123, 3, 22, 11, 7, 119, 120, 10, 5, 0, 0, 120, 121, 7, 2, 0, 0, 121, 123,
		3, 22, 11, 6, 122, 113, 1, 0, 0, 0, 122, 116, 1, 0, 0, 0, 122, 119, 1,
		0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0,
		0, 125, 23, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 7, 27, 41, 85, 97, 111, 122,
		124,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TopLangParserInit initializes any static state used to implement TopLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTopLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TopLangParserInit() {
	staticData := &TopLangParserStaticData
	staticData.once.Do(toplangParserInit)
}

// NewTopLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTopLangParser(input antlr.TokenStream) *TopLangParser {
	TopLangParserInit()
	this := new(TopLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TopLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TopLang.g4"

	return this
}

// TopLangParser tokens.
const (
	TopLangParserEOF         = antlr.TokenEOF
	TopLangParserVAR         = 1
	TopLangParserPRINT       = 2
	TopLangParserPRINTNUMBER = 3
	TopLangParserPRINTSTRING = 4
	TopLangParserREAD        = 5
	TopLangParserIF          = 6
	TopLangParserELSE        = 7
	TopLangParserWHILE       = 8
	TopLangParserEQ          = 9
	TopLangParserNEQ         = 10
	TopLangParserLTE         = 11
	TopLangParserGTE         = 12
	TopLangParserLT          = 13
	TopLangParserGT          = 14
	TopLangParserPLUS        = 15
	TopLangParserMINUS       = 16
	TopLangParserMUL         = 17
	TopLangParserDIV         = 18
	TopLangParserMOD         = 19
	TopLangParserASSIGN      = 20
	TopLangParserLPAREN      = 21
	TopLangParserRPAREN      = 22
	TopLangParserLBRACE      = 23
	TopLangParserRBRACE      = 24
	TopLangParserSEMI        = 25
	TopLangParserINT         = 26
	TopLangParserSTRING      = 27
	TopLangParserID          = 28
	TopLangParserWS          = 29
	TopLangParserCOMMENT     = 30
)

// TopLangParser rules.
const (
	TopLangParserRULE_program         = 0
	TopLangParserRULE_statement       = 1
	TopLangParserRULE_varDecl         = 2
	TopLangParserRULE_assignment      = 3
	TopLangParserRULE_printStmt       = 4
	TopLangParserRULE_printNumberStmt = 5
	TopLangParserRULE_printStringStmt = 6
	TopLangParserRULE_readStmt        = 7
	TopLangParserRULE_ifStmt          = 8
	TopLangParserRULE_whileStmt       = 9
	TopLangParserRULE_block           = 10
	TopLangParserRULE_expr            = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(TopLangParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TopLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&276824446) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(TopLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	Assignment() IAssignmentContext
	PrintStmt() IPrintStmtContext
	PrintNumberStmt() IPrintNumberStmtContext
	PrintStringStmt() IPrintStringStmtContext
	ReadStmt() IReadStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	Block() IBlockContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementContext) PrintNumberStmt() IPrintNumberStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintNumberStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintNumberStmtContext)
}

func (s *StatementContext) PrintStringStmt() IPrintStringStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStringStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStringStmtContext)
}

func (s *StatementContext) ReadStmt() IReadStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TopLangParserRULE_statement)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TopLangParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.VarDecl()
		}

	case TopLangParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Assignment()
		}

	case TopLangParserPRINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.PrintStmt()
		}

	case TopLangParserPRINTNUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
			p.PrintNumberStmt()
		}

	case TopLangParserPRINTSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(36)
			p.PrintStringStmt()
		}

	case TopLangParserREAD:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(37)
			p.ReadStmt()
		}

	case TopLangParserIF:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(38)
			p.IfStmt()
		}

	case TopLangParserWHILE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(39)
			p.WhileStmt()
		}

	case TopLangParserLBRACE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(40)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SEMI() antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(TopLangParserVAR, 0)
}

func (s *VarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TopLangParserID, 0)
}

func (s *VarDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(TopLangParserASSIGN, 0)
}

func (s *VarDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TopLangParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(TopLangParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(TopLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(TopLangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.expr(0)
	}
	{
		p.SetState(47)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SEMI() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(TopLangParserID, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(TopLangParserASSIGN, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TopLangParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(TopLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(TopLangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.expr(0)
	}
	{
		p.SetState(52)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(TopLangParserPRINT, 0)
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *PrintStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *PrintStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TopLangParserRULE_printStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(TopLangParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.expr(0)
	}
	{
		p.SetState(57)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintNumberStmtContext is an interface to support dynamic dispatch.
type IPrintNumberStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINTNUMBER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsPrintNumberStmtContext differentiates from other interfaces.
	IsPrintNumberStmtContext()
}

type PrintNumberStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintNumberStmtContext() *PrintNumberStmtContext {
	var p = new(PrintNumberStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printNumberStmt
	return p
}

func InitEmptyPrintNumberStmtContext(p *PrintNumberStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printNumberStmt
}

func (*PrintNumberStmtContext) IsPrintNumberStmtContext() {}

func NewPrintNumberStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintNumberStmtContext {
	var p = new(PrintNumberStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_printNumberStmt

	return p
}

func (s *PrintNumberStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintNumberStmtContext) PRINTNUMBER() antlr.TerminalNode {
	return s.GetToken(TopLangParserPRINTNUMBER, 0)
}

func (s *PrintNumberStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *PrintNumberStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintNumberStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *PrintNumberStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *PrintNumberStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintNumberStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintNumberStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterPrintNumberStmt(s)
	}
}

func (s *PrintNumberStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitPrintNumberStmt(s)
	}
}

func (s *PrintNumberStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitPrintNumberStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) PrintNumberStmt() (localctx IPrintNumberStmtContext) {
	localctx = NewPrintNumberStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TopLangParserRULE_printNumberStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(TopLangParserPRINTNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.expr(0)
	}
	{
		p.SetState(63)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStringStmtContext is an interface to support dynamic dispatch.
type IPrintStringStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINTSTRING() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsPrintStringStmtContext differentiates from other interfaces.
	IsPrintStringStmtContext()
}

type PrintStringStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStringStmtContext() *PrintStringStmtContext {
	var p = new(PrintStringStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printStringStmt
	return p
}

func InitEmptyPrintStringStmtContext(p *PrintStringStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_printStringStmt
}

func (*PrintStringStmtContext) IsPrintStringStmtContext() {}

func NewPrintStringStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStringStmtContext {
	var p = new(PrintStringStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_printStringStmt

	return p
}

func (s *PrintStringStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStringStmtContext) PRINTSTRING() antlr.TerminalNode {
	return s.GetToken(TopLangParserPRINTSTRING, 0)
}

func (s *PrintStringStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *PrintStringStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(TopLangParserSTRING, 0)
}

func (s *PrintStringStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *PrintStringStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *PrintStringStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStringStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStringStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterPrintStringStmt(s)
	}
}

func (s *PrintStringStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitPrintStringStmt(s)
	}
}

func (s *PrintStringStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitPrintStringStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) PrintStringStmt() (localctx IPrintStringStmtContext) {
	localctx = NewPrintStringStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TopLangParserRULE_printStringStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(TopLangParserPRINTSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(TopLangParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadStmtContext is an interface to support dynamic dispatch.
type IReadStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ID() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsReadStmtContext differentiates from other interfaces.
	IsReadStmtContext()
}

type ReadStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadStmtContext() *ReadStmtContext {
	var p = new(ReadStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_readStmt
	return p
}

func InitEmptyReadStmtContext(p *ReadStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_readStmt
}

func (*ReadStmtContext) IsReadStmtContext() {}

func NewReadStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadStmtContext {
	var p = new(ReadStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_readStmt

	return p
}

func (s *ReadStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadStmtContext) READ() antlr.TerminalNode {
	return s.GetToken(TopLangParserREAD, 0)
}

func (s *ReadStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *ReadStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(TopLangParserID, 0)
}

func (s *ReadStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *ReadStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(TopLangParserSEMI, 0)
}

func (s *ReadStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterReadStmt(s)
	}
}

func (s *ReadStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitReadStmt(s)
	}
}

func (s *ReadStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitReadStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) ReadStmt() (localctx IReadStmtContext) {
	localctx = NewReadStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TopLangParserRULE_readStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(TopLangParserREAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(TopLangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(TopLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(TopLangParserIF, 0)
}

func (s *IfStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TopLangParserELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TopLangParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(TopLangParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.expr(0)
	}
	{
		p.SetState(81)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Block()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TopLangParserELSE {
		{
			p.SetState(83)
			p.Match(TopLangParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(TopLangParserWHILE, 0)
}

func (s *WhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TopLangParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(TopLangParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(TopLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.expr(0)
	}
	{
		p.SetState(90)
		p.Match(TopLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TopLangParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TopLangParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TopLangParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(TopLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&276824446) != 0 {
		{
			p.SetState(94)
			p.Statement()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(TopLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TopLangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TopLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulExprContext) GetLeft() IExprContext { return s.left }

func (s *MulExprContext) GetRight() IExprContext { return s.right }

func (s *MulExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *MulExprContext) SetRight(v IExprContext) { s.right = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(TopLangParserMUL, 0)
}

func (s *MulExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(TopLangParserDIV, 0)
}

func (s *MulExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(TopLangParserMOD, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(TopLangParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	ExprContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(TopLangParserINT, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewCompExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExprContext {
	var p = new(CompExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CompExprContext) GetOp() antlr.Token { return s.op }

func (s *CompExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompExprContext) GetLeft() IExprContext { return s.left }

func (s *CompExprContext) GetRight() IExprContext { return s.right }

func (s *CompExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *CompExprContext) SetRight(v IExprContext) { s.right = v }

func (s *CompExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(TopLangParserEQ, 0)
}

func (s *CompExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(TopLangParserNEQ, 0)
}

func (s *CompExprContext) LT() antlr.TerminalNode {
	return s.GetToken(TopLangParserLT, 0)
}

func (s *CompExprContext) GT() antlr.TerminalNode {
	return s.GetToken(TopLangParserGT, 0)
}

func (s *CompExprContext) LTE() antlr.TerminalNode {
	return s.GetToken(TopLangParserLTE, 0)
}

func (s *CompExprContext) GTE() antlr.TerminalNode {
	return s.GetToken(TopLangParserGTE, 0)
}

func (s *CompExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterCompExpr(s)
	}
}

func (s *CompExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitCompExpr(s)
	}
}

func (s *CompExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitCompExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExprContext) GetLeft() IExprContext { return s.left }

func (s *AddExprContext) GetRight() IExprContext { return s.right }

func (s *AddExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddExprContext) SetRight(v IExprContext) { s.right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TopLangParserPLUS, 0)
}

func (s *AddExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TopLangParserMINUS, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TopLangParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExprContext struct {
	ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TopLangParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TopLangListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TopLangVisitor:
		return t.VisitUnaryMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TopLangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TopLangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, TopLangParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TopLangParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(103)
			p.Match(TopLangParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.expr(4)
		}

	case TopLangParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.Match(TopLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.expr(0)
		}
		{
			p.SetState(107)
			p.Match(TopLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TopLangParserINT:
		localctx = NewIntLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.Match(TopLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TopLangParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.Match(TopLangParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCompExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CompExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TopLangParserRULE_expr)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(114)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)

					var _x = p.expr(8)

					localctx.(*CompExprContext).right = _x
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TopLangParserRULE_expr)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(117)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TopLangParserPLUS || _la == TopLangParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)

					var _x = p.expr(7)

					localctx.(*AddExprContext).right = _x
				}

			case 3:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TopLangParserRULE_expr)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&917504) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(121)

					var _x = p.expr(6)

					localctx.(*MulExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TopLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TopLangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
