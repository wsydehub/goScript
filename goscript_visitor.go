// Code generated from GoScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GoScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GoScriptParser.
type GoScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoScriptParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by GoScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by GoScriptParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by GoScriptParser#formalParameterDecl.
	VisitFormalParameterDecl(ctx *FormalParameterDeclContext) interface{}

	// Visit a parse tree produced by GoScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GoScriptParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by GoScriptParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by GoScriptParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by GoScriptParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by GoScriptParser#arrayInitializerElement.
	VisitArrayInitializerElement(ctx *ArrayInitializerElementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#mapInitializer.
	VisitMapInitializer(ctx *MapInitializerContext) interface{}

	// Visit a parse tree produced by GoScriptParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by GoScriptParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by GoScriptParser#connectorType.
	VisitConnectorType(ctx *ConnectorTypeContext) interface{}

	// Visit a parse tree produced by GoScriptParser#dynamicType.
	VisitDynamicType(ctx *DynamicTypeContext) interface{}

	// Visit a parse tree produced by GoScriptParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CommonBlockStatement.
	VisitCommonBlockStatement(ctx *CommonBlockStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#IfStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ExprStatement.
	VisitExprStatement(ctx *ExprStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CreatorExpr.
	VisitCreatorExpr(ctx *CreatorExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#FetchExpr.
	VisitFetchExpr(ctx *FetchExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#IndexExpr.
	VisitIndexExpr(ctx *IndexExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CreatorAndAssignExpr.
	VisitCreatorAndAssignExpr(ctx *CreatorAndAssignExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AssignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#NegativeExpr.
	VisitNegativeExpr(ctx *NegativeExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CompareExpr.
	VisitCompareExpr(ctx *CompareExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#EqNEqExpr.
	VisitEqNEqExpr(ctx *EqNEqExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#PrimaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#TernaryExpr.
	VisitTernaryExpr(ctx *TernaryExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by GoScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GoScriptParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by GoScriptParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by GoScriptParser#MapCreator.
	VisitMapCreator(ctx *MapCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ArrayCreator.
	VisitArrayCreator(ctx *ArrayCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#MeegoCreator.
	VisitMeegoCreator(ctx *MeegoCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#PrimitiveCreator.
	VisitPrimitiveCreator(ctx *PrimitiveCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#DynamicCreator.
	VisitDynamicCreator(ctx *DynamicCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#creatorName.
	VisitCreatorName(ctx *CreatorNameContext) interface{}
}
