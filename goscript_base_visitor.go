// Code generated from GoScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GoScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGoScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoScriptVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitFormalParameterDecl(ctx *FormalParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitArrayInitializerElement(ctx *ArrayInitializerElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitMapInitializer(ctx *MapInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitConnectorType(ctx *ConnectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitDynamicType(ctx *DynamicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCommonBlockStatement(ctx *CommonBlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitExprStatement(ctx *ExprStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitForControl(ctx *ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCreatorExpr(ctx *CreatorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitFetchExpr(ctx *FetchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCreatorAndAssignExpr(ctx *CreatorAndAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitNegativeExpr(ctx *NegativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCompareExpr(ctx *CompareExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitEqNEqExpr(ctx *EqNEqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitMapCreator(ctx *MapCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitArrayCreator(ctx *ArrayCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitMeegoCreator(ctx *MeegoCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitPrimitiveCreator(ctx *PrimitiveCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitDynamicCreator(ctx *DynamicCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoScriptVisitor) VisitCreatorName(ctx *CreatorNameContext) interface{} {
	return v.VisitChildren(ctx)
}
