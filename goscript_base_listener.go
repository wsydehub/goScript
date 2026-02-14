// Code generated from GoScript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package goScript // GoScript
import "github.com/antlr4-go/antlr/v4"

// BaseGoScriptListener is a complete listener for a parse tree produced by GoScriptParser.
type BaseGoScriptListener struct{}

var _ GoScriptListener = &BaseGoScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseGoScriptListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseGoScriptListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseGoScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseGoScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseGoScriptListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseGoScriptListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterDecl is called when production formalParameterDecl is entered.
func (s *BaseGoScriptListener) EnterFormalParameterDecl(ctx *FormalParameterDeclContext) {}

// ExitFormalParameterDecl is called when production formalParameterDecl is exited.
func (s *BaseGoScriptListener) ExitFormalParameterDecl(ctx *FormalParameterDeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseGoScriptListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseGoScriptListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseGoScriptListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseGoScriptListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseGoScriptListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseGoScriptListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseGoScriptListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseGoScriptListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseGoScriptListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseGoScriptListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseGoScriptListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseGoScriptListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterMapInitializer is called when production mapInitializer is entered.
func (s *BaseGoScriptListener) EnterMapInitializer(ctx *MapInitializerContext) {}

// ExitMapInitializer is called when production mapInitializer is exited.
func (s *BaseGoScriptListener) ExitMapInitializer(ctx *MapInitializerContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseGoScriptListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseGoScriptListener) ExitType_(ctx *Type_Context) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseGoScriptListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseGoScriptListener) ExitMapType(ctx *MapTypeContext) {}

// EnterConnectorType is called when production connectorType is entered.
func (s *BaseGoScriptListener) EnterConnectorType(ctx *ConnectorTypeContext) {}

// ExitConnectorType is called when production connectorType is exited.
func (s *BaseGoScriptListener) ExitConnectorType(ctx *ConnectorTypeContext) {}

// EnterDynamicType is called when production dynamicType is entered.
func (s *BaseGoScriptListener) EnterDynamicType(ctx *DynamicTypeContext) {}

// ExitDynamicType is called when production dynamicType is exited.
func (s *BaseGoScriptListener) ExitDynamicType(ctx *DynamicTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseGoScriptListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseGoScriptListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseGoScriptListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseGoScriptListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseGoScriptListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseGoScriptListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseGoScriptListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseGoScriptListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseGoScriptListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseGoScriptListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseGoScriptListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseGoScriptListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseGoScriptListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseGoScriptListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseGoScriptListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseGoScriptListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseGoScriptListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseGoScriptListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseGoScriptListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseGoScriptListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseGoScriptListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseGoScriptListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseGoScriptListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseGoScriptListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterCreateAndAssignExpr is called when production CreateAndAssignExpr is entered.
func (s *BaseGoScriptListener) EnterCreateAndAssignExpr(ctx *CreateAndAssignExprContext) {}

// ExitCreateAndAssignExpr is called when production CreateAndAssignExpr is exited.
func (s *BaseGoScriptListener) ExitCreateAndAssignExpr(ctx *CreateAndAssignExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseGoScriptListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseGoScriptListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterConditionalExpr is called when production ConditionalExpr is entered.
func (s *BaseGoScriptListener) EnterConditionalExpr(ctx *ConditionalExprContext) {}

// ExitConditionalExpr is called when production ConditionalExpr is exited.
func (s *BaseGoScriptListener) ExitConditionalExpr(ctx *ConditionalExprContext) {}

// EnterUnaryExpr is called when production UnaryExpr is entered.
func (s *BaseGoScriptListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production UnaryExpr is exited.
func (s *BaseGoScriptListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseGoScriptListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseGoScriptListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterIndexExpr is called when production IndexExpr is entered.
func (s *BaseGoScriptListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production IndexExpr is exited.
func (s *BaseGoScriptListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterAssignExpr is called when production AssignExpr is entered.
func (s *BaseGoScriptListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production AssignExpr is exited.
func (s *BaseGoScriptListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterSelectorExpr is called when production SelectorExpr is entered.
func (s *BaseGoScriptListener) EnterSelectorExpr(ctx *SelectorExprContext) {}

// ExitSelectorExpr is called when production SelectorExpr is exited.
func (s *BaseGoScriptListener) ExitSelectorExpr(ctx *SelectorExprContext) {}

// EnterCreateExpr is called when production CreateExpr is entered.
func (s *BaseGoScriptListener) EnterCreateExpr(ctx *CreateExprContext) {}

// ExitCreateExpr is called when production CreateExpr is exited.
func (s *BaseGoScriptListener) ExitCreateExpr(ctx *CreateExprContext) {}

// EnterSelfAddExpr is called when production SelfAddExpr is entered.
func (s *BaseGoScriptListener) EnterSelfAddExpr(ctx *SelfAddExprContext) {}

// ExitSelfAddExpr is called when production SelfAddExpr is exited.
func (s *BaseGoScriptListener) ExitSelfAddExpr(ctx *SelfAddExprContext) {}

// EnterPrimaryExpr is called when production PrimaryExpr is entered.
func (s *BaseGoScriptListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production PrimaryExpr is exited.
func (s *BaseGoScriptListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseGoScriptListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseGoScriptListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterTernaryExpr is called when production TernaryExpr is entered.
func (s *BaseGoScriptListener) EnterTernaryExpr(ctx *TernaryExprContext) {}

// ExitTernaryExpr is called when production TernaryExpr is exited.
func (s *BaseGoScriptListener) ExitTernaryExpr(ctx *TernaryExprContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseGoScriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseGoScriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseGoScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseGoScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseGoScriptListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseGoScriptListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGoScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGoScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseGoScriptListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseGoScriptListener) ExitCreator(ctx *CreatorContext) {}

// EnterMapCreator is called when production mapCreator is entered.
func (s *BaseGoScriptListener) EnterMapCreator(ctx *MapCreatorContext) {}

// ExitMapCreator is called when production mapCreator is exited.
func (s *BaseGoScriptListener) ExitMapCreator(ctx *MapCreatorContext) {}

// EnterArrayCreator is called when production arrayCreator is entered.
func (s *BaseGoScriptListener) EnterArrayCreator(ctx *ArrayCreatorContext) {}

// ExitArrayCreator is called when production arrayCreator is exited.
func (s *BaseGoScriptListener) ExitArrayCreator(ctx *ArrayCreatorContext) {}

// EnterCreatorName is called when production creatorName is entered.
func (s *BaseGoScriptListener) EnterCreatorName(ctx *CreatorNameContext) {}

// ExitCreatorName is called when production creatorName is exited.
func (s *BaseGoScriptListener) ExitCreatorName(ctx *CreatorNameContext) {}

// EnterConnectorCreator is called when production connectorCreator is entered.
func (s *BaseGoScriptListener) EnterConnectorCreator(ctx *ConnectorCreatorContext) {}

// ExitConnectorCreator is called when production connectorCreator is exited.
func (s *BaseGoScriptListener) ExitConnectorCreator(ctx *ConnectorCreatorContext) {}

// EnterPrimitiveCreator is called when production primitiveCreator is entered.
func (s *BaseGoScriptListener) EnterPrimitiveCreator(ctx *PrimitiveCreatorContext) {}

// ExitPrimitiveCreator is called when production primitiveCreator is exited.
func (s *BaseGoScriptListener) ExitPrimitiveCreator(ctx *PrimitiveCreatorContext) {}

// EnterDynamicCreator is called when production dynamicCreator is entered.
func (s *BaseGoScriptListener) EnterDynamicCreator(ctx *DynamicCreatorContext) {}

// ExitDynamicCreator is called when production dynamicCreator is exited.
func (s *BaseGoScriptListener) ExitDynamicCreator(ctx *DynamicCreatorContext) {}
