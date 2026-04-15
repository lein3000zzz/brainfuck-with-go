// Code generated from tools/TopLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TopLang
import "github.com/antlr4-go/antlr/v4"

type BaseTopLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTopLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitPrintNumberStmt(ctx *PrintNumberStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitPrintStringStmt(ctx *PrintStringStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitReadStmt(ctx *ReadStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitCompExpr(ctx *CompExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTopLangVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}
