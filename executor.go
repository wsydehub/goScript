package goScript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
)

type Executor struct {
	breakFlag    bool
	continueFlag bool
	returnFlag   bool
	killFlag     bool

	scopeStack *ScopeStack
	funcStack  *FuncStack
	funcMap    map[string]*FuncStack
	varMap     map[string]*Variable
	GoVarMap   map[string]reflect.Value
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) PushFunc(f *Function) {
	e.funcStack.Push(f)
	e.scopeStack.Push(f.Scope)
}

func (e *Executor) PopFunc() *Function {
	e.scopeStack.Pop()
	return e.funcStack.Pop()
}

func (e *Executor) PushScope(s *Scope) {
	e.scopeStack.Push(s)
}

func (e *Executor) PopScope() *Scope {
	return e.scopeStack.Pop()
}

func (e *Executor) Visit(tree antlr.ParseTree) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitFormalParameterDecl(ctx *FormalParameterDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitBlock(ctx *BlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitMapInitializer(ctx *MapInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitType_(ctx *Type_Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitMapType(ctx *MapTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitConnectorType(ctx *ConnectorTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitDynamicType(ctx *DynamicTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitStatement(ctx *StatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitForStatement(ctx *ForStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitForControl(ctx *ForControlContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitForInit(ctx *ForInitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitMulExpr(ctx *MulExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitAndExpr(ctx *AndExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCreateAndAssignExpr(ctx *CreateAndAssignExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitAddExpr(ctx *AddExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitConditionalExpr(ctx *ConditionalExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitOrExpr(ctx *OrExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCreateExpr(ctx *CreateExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitSelfAddExpr(ctx *SelfAddExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCallExpr(ctx *CallExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitPrimary(ctx *PrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitLiteral(ctx *LiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCreator(ctx *CreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitMapCreator(ctx *MapCreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitArrayCreator(ctx *ArrayCreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitCreatorName(ctx *CreatorNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitConnectorCreator(ctx *ConnectorCreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitPrimitiveCreator(ctx *PrimitiveCreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (e *Executor) VisitDynamicCreator(ctx *DynamicCreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}
