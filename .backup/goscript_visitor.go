// Code generated from GoScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package goScript // GoScript

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

	// Visit a parse tree produced by GoScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GoScriptParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by GoScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by GoScriptParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CreateAndAssignExpr.
	VisitCreateAndAssignExpr(ctx *CreateAndAssignExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#ConditionalExpr.
	VisitConditionalExpr(ctx *ConditionalExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#IndexExpr.
	VisitIndexExpr(ctx *IndexExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#AssignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#SelectorExpr.
	VisitSelectorExpr(ctx *SelectorExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CreateExpr.
	VisitCreateExpr(ctx *CreateExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#SelfAddExpr.
	VisitSelfAddExpr(ctx *SelfAddExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#PrimaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by GoScriptParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

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

	// Visit a parse tree produced by GoScriptParser#creator.
	VisitCreator(ctx *CreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#mapCreator.
	VisitMapCreator(ctx *MapCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#arrayCreator.
	VisitArrayCreator(ctx *ArrayCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#creatorName.
	VisitCreatorName(ctx *CreatorNameContext) interface{}

	// Visit a parse tree produced by GoScriptParser#connectorCreator.
	VisitConnectorCreator(ctx *ConnectorCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#primitiveCreator.
	VisitPrimitiveCreator(ctx *PrimitiveCreatorContext) interface{}

	// Visit a parse tree produced by GoScriptParser#dynamicCreator.
	VisitDynamicCreator(ctx *DynamicCreatorContext) interface{}
}
