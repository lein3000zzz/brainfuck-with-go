// Code generated from tools/TopLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TopLang
import "github.com/antlr4-go/antlr/v4"

// BaseTopLangListener is a complete listener for a parse tree produced by TopLangParser.
type BaseTopLangListener struct{}

var _ TopLangListener = &BaseTopLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTopLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTopLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTopLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTopLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseTopLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTopLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTopLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTopLangListener) ExitStatement(ctx *StatementContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseTopLangListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseTopLangListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseTopLangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseTopLangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseTopLangListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseTopLangListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterPrintNumberStmt is called when production printNumberStmt is entered.
func (s *BaseTopLangListener) EnterPrintNumberStmt(ctx *PrintNumberStmtContext) {}

// ExitPrintNumberStmt is called when production printNumberStmt is exited.
func (s *BaseTopLangListener) ExitPrintNumberStmt(ctx *PrintNumberStmtContext) {}

// EnterPrintStringStmt is called when production printStringStmt is entered.
func (s *BaseTopLangListener) EnterPrintStringStmt(ctx *PrintStringStmtContext) {}

// ExitPrintStringStmt is called when production printStringStmt is exited.
func (s *BaseTopLangListener) ExitPrintStringStmt(ctx *PrintStringStmtContext) {}

// EnterReadStmt is called when production readStmt is entered.
func (s *BaseTopLangListener) EnterReadStmt(ctx *ReadStmtContext) {}

// ExitReadStmt is called when production readStmt is exited.
func (s *BaseTopLangListener) ExitReadStmt(ctx *ReadStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseTopLangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseTopLangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseTopLangListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseTopLangListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTopLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTopLangListener) ExitBlock(ctx *BlockContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseTopLangListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseTopLangListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseTopLangListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseTopLangListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseTopLangListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseTopLangListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterCompExpr is called when production CompExpr is entered.
func (s *BaseTopLangListener) EnterCompExpr(ctx *CompExprContext) {}

// ExitCompExpr is called when production CompExpr is exited.
func (s *BaseTopLangListener) ExitCompExpr(ctx *CompExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseTopLangListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseTopLangListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseTopLangListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseTopLangListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterUnaryMinusExpr is called when production UnaryMinusExpr is entered.
func (s *BaseTopLangListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production UnaryMinusExpr is exited.
func (s *BaseTopLangListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}
