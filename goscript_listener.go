// Code generated from GoScript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package goScript // GoScript
import "github.com/antlr4-go/antlr/v4"

// GoScriptListener is a complete listener for a parse tree produced by GoScriptParser.
type GoScriptListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterTopLevelStatement is called when entering the topLevelStatement production.
	EnterTopLevelStatement(c *TopLevelStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterDecl is called when entering the formalParameterDecl production.
	EnterFormalParameterDecl(c *FormalParameterDeclContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterMapInitializer is called when entering the mapInitializer production.
	EnterMapInitializer(c *MapInitializerContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterConnectorType is called when entering the connectorType production.
	EnterConnectorType(c *ConnectorTypeContext)

	// EnterDynamicType is called when entering the dynamicType production.
	EnterDynamicType(c *DynamicTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterCreateAndAssignExpr is called when entering the CreateAndAssignExpr production.
	EnterCreateAndAssignExpr(c *CreateAndAssignExprContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterConditionalExpr is called when entering the ConditionalExpr production.
	EnterConditionalExpr(c *ConditionalExprContext)

	// EnterUnaryExpr is called when entering the UnaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterAssignExpr is called when entering the AssignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterSelectorExpr is called when entering the SelectorExpr production.
	EnterSelectorExpr(c *SelectorExprContext)

	// EnterCreateExpr is called when entering the CreateExpr production.
	EnterCreateExpr(c *CreateExprContext)

	// EnterSelfAddExpr is called when entering the SelfAddExpr production.
	EnterSelfAddExpr(c *SelfAddExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterTernaryExpr is called when entering the TernaryExpr production.
	EnterTernaryExpr(c *TernaryExprContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterMapCreator is called when entering the mapCreator production.
	EnterMapCreator(c *MapCreatorContext)

	// EnterArrayCreator is called when entering the arrayCreator production.
	EnterArrayCreator(c *ArrayCreatorContext)

	// EnterCreatorName is called when entering the creatorName production.
	EnterCreatorName(c *CreatorNameContext)

	// EnterConnectorCreator is called when entering the connectorCreator production.
	EnterConnectorCreator(c *ConnectorCreatorContext)

	// EnterPrimitiveCreator is called when entering the primitiveCreator production.
	EnterPrimitiveCreator(c *PrimitiveCreatorContext)

	// EnterDynamicCreator is called when entering the dynamicCreator production.
	EnterDynamicCreator(c *DynamicCreatorContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitTopLevelStatement is called when exiting the topLevelStatement production.
	ExitTopLevelStatement(c *TopLevelStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterDecl is called when exiting the formalParameterDecl production.
	ExitFormalParameterDecl(c *FormalParameterDeclContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitMapInitializer is called when exiting the mapInitializer production.
	ExitMapInitializer(c *MapInitializerContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitConnectorType is called when exiting the connectorType production.
	ExitConnectorType(c *ConnectorTypeContext)

	// ExitDynamicType is called when exiting the dynamicType production.
	ExitDynamicType(c *DynamicTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitCreateAndAssignExpr is called when exiting the CreateAndAssignExpr production.
	ExitCreateAndAssignExpr(c *CreateAndAssignExprContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitConditionalExpr is called when exiting the ConditionalExpr production.
	ExitConditionalExpr(c *ConditionalExprContext)

	// ExitUnaryExpr is called when exiting the UnaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitAssignExpr is called when exiting the AssignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitSelectorExpr is called when exiting the SelectorExpr production.
	ExitSelectorExpr(c *SelectorExprContext)

	// ExitCreateExpr is called when exiting the CreateExpr production.
	ExitCreateExpr(c *CreateExprContext)

	// ExitSelfAddExpr is called when exiting the SelfAddExpr production.
	ExitSelfAddExpr(c *SelfAddExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitTernaryExpr is called when exiting the TernaryExpr production.
	ExitTernaryExpr(c *TernaryExprContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitMapCreator is called when exiting the mapCreator production.
	ExitMapCreator(c *MapCreatorContext)

	// ExitArrayCreator is called when exiting the arrayCreator production.
	ExitArrayCreator(c *ArrayCreatorContext)

	// ExitCreatorName is called when exiting the creatorName production.
	ExitCreatorName(c *CreatorNameContext)

	// ExitConnectorCreator is called when exiting the connectorCreator production.
	ExitConnectorCreator(c *ConnectorCreatorContext)

	// ExitPrimitiveCreator is called when exiting the primitiveCreator production.
	ExitPrimitiveCreator(c *PrimitiveCreatorContext)

	// ExitDynamicCreator is called when exiting the dynamicCreator production.
	ExitDynamicCreator(c *DynamicCreatorContext)
}
