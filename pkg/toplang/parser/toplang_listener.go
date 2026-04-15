// Code generated from tools/TopLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TopLang
import "github.com/antlr4-go/antlr/v4"

// TopLangListener is a complete listener for a parse tree produced by TopLangParser.
type TopLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterPrintNumberStmt is called when entering the printNumberStmt production.
	EnterPrintNumberStmt(c *PrintNumberStmtContext)

	// EnterPrintStringStmt is called when entering the printStringStmt production.
	EnterPrintStringStmt(c *PrintStringStmtContext)

	// EnterReadStmt is called when entering the readStmt production.
	EnterReadStmt(c *ReadStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterCompExpr is called when entering the CompExpr production.
	EnterCompExpr(c *CompExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterUnaryMinusExpr is called when entering the UnaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitPrintNumberStmt is called when exiting the printNumberStmt production.
	ExitPrintNumberStmt(c *PrintNumberStmtContext)

	// ExitPrintStringStmt is called when exiting the printStringStmt production.
	ExitPrintStringStmt(c *PrintStringStmtContext)

	// ExitReadStmt is called when exiting the readStmt production.
	ExitReadStmt(c *ReadStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitCompExpr is called when exiting the CompExpr production.
	ExitCompExpr(c *CompExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitUnaryMinusExpr is called when exiting the UnaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)
}
