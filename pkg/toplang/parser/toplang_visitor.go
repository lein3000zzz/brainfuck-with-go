// Code generated from tools/TopLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TopLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TopLangParser.
type TopLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TopLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TopLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TopLangParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by TopLangParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by TopLangParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#printNumberStmt.
	VisitPrintNumberStmt(ctx *PrintNumberStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#printStringStmt.
	VisitPrintStringStmt(ctx *PrintStringStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#readStmt.
	VisitReadStmt(ctx *ReadStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by TopLangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by TopLangParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by TopLangParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by TopLangParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by TopLangParser#CompExpr.
	VisitCompExpr(ctx *CompExprContext) interface{}

	// Visit a parse tree produced by TopLangParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by TopLangParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by TopLangParser#UnaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}
}
