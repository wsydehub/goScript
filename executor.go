package goScript

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
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
	executor := &Executor{
		scopeStack: NewScopeStack(),
		funcStack:  NewFuncStack(),
		funcMap:    map[string]*FuncStack{},
		varMap:     map[string]*Variable{},
		GoVarMap:   map[string]reflect.Value{},
	}
	executor.scopeStack.Push(NewScope(CommonScopeType))
	return executor
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
	return tree.Accept(e)
}

func (e *Executor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if childTree, ok := child.(antlr.ParseTree); ok {
			result = childTree.Accept(e)
		}
	}
	return result
}

func (e *Executor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}

func (e *Executor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (e *Executor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return nil
}

func (e *Executor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return nil
}

func (e *Executor) VisitFormalParameterDecl(ctx *FormalParameterDeclContext) interface{} {
	return nil
}

func (e *Executor) VisitBlock(ctx *BlockContext) interface{} {
	e.PushScope(NewScope(CommonScopeType))
	defer e.PopScope()
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	varType, typeGo := e.typeFromText(ctx.Type_().GetText())
	for _, decl := range ctx.VariableDeclarators().AllVariableDeclarator() {
		declCtx := decl.(*VariableDeclaratorContext)
		name := declCtx.Identifier().GetText()
		value := e.defaultValue(varType)
		if declCtx.VariableInitializer() != nil {
			value = declCtx.VariableInitializer().Accept(e)
		}
		coerced := e.coerceValue(value, varType)
		var valueRef reflect.Value
		if coerced == nil {
			valueRef = reflect.Zero(typeGo)
		} else {
			valueRef = reflect.ValueOf(coerced)
			if typeGo != nil && valueRef.IsValid() && valueRef.Type() != typeGo && valueRef.Type().ConvertibleTo(typeGo) {
				valueRef = valueRef.Convert(typeGo)
			}
		}
		e.addVar(NewVariable(name, varType, typeGo, valueRef))
	}
	return nil
}

func (e *Executor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return nil
}

func (e *Executor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(e)
	}
	return nil
}

func (e *Executor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return nil
}

func (e *Executor) VisitMapInitializer(ctx *MapInitializerContext) interface{} {
	return nil
}

func (e *Executor) VisitType_(ctx *Type_Context) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitMapType(ctx *MapTypeContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitConnectorType(ctx *ConnectorTypeContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitDynamicType(ctx *DynamicTypeContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitStatement(ctx *StatementContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	cond := e.toBool(ctx.Expression().Accept(e))
	if cond {
		return ctx.Statement(0).Accept(e)
	}
	if ctx.Statement(1) != nil {
		return ctx.Statement(1).Accept(e)
	}
	return nil
}

func (e *Executor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return nil
}

func (e *Executor) VisitForControl(ctx *ForControlContext) interface{} {
	return nil
}

func (e *Executor) VisitForInit(ctx *ForInitContext) interface{} {
	return nil
}

func (e *Executor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return nil
}

func (e *Executor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(e)
	}
	return nil
}

func (e *Executor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return nil
}

func (e *Executor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return nil
}

func (e *Executor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return ctx.Expression().Accept(e)
}

func (e *Executor) VisitMulExpr(ctx *MulExprContext) interface{} {
	left := ctx.Expression(0).Accept(e)
	right := ctx.Expression(1).Accept(e)
	if ctx.Multiplication() != nil {
		return e.mul(left, right)
	}
	if ctx.Division() != nil {
		return e.div(left, right)
	}
	return e.mod(left, right)
}

func (e *Executor) VisitAndExpr(ctx *AndExprContext) interface{} {
	left := e.toBool(ctx.Expression(0).Accept(e))
	if !left {
		return false
	}
	return e.toBool(ctx.Expression(1).Accept(e))
}

func (e *Executor) VisitCreateAndAssignExpr(ctx *CreateAndAssignExprContext) interface{} {
	values := ctx.AllExpression()
	if len(values) == 0 {
		return nil
	}
	return values[len(values)-1].Accept(e)
}

func (e *Executor) VisitAddExpr(ctx *AddExprContext) interface{} {
	left := ctx.Expression(0).Accept(e)
	right := ctx.Expression(1).Accept(e)
	if ctx.Plus() != nil {
		return e.add(left, right)
	}
	return e.sub(left, right)
}

func (e *Executor) VisitConditionalExpr(ctx *ConditionalExprContext) interface{} {
	left := ctx.Expression(0).Accept(e)
	right := ctx.Expression(1).Accept(e)
	if ctx.Let() != nil {
		return e.compare(left, right, "<=")
	}
	if ctx.Get() != nil {
		return e.compare(left, right, ">=")
	}
	if ctx.Gt() != nil {
		return e.compare(left, right, ">")
	}
	if ctx.Lt() != nil {
		return e.compare(left, right, "<")
	}
	if ctx.Eq() != nil {
		return e.compare(left, right, "==")
	}
	return e.compare(left, right, "!=")
}

func (e *Executor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	value := ctx.Expression().Accept(e)
	if ctx.Plus() != nil {
		return value
	}
	if ctx.Minus() != nil {
		return e.negate(value)
	}
	return !e.toBool(value)
}

func (e *Executor) VisitOrExpr(ctx *OrExprContext) interface{} {
	left := e.toBool(ctx.Expression(0).Accept(e))
	if left {
		return true
	}
	return e.toBool(ctx.Expression(1).Accept(e))
}

func (e *Executor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	return nil
}

func (e *Executor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	values := ctx.AllExpression()
	if len(values) == 0 {
		return nil
	}
	return values[len(values)-1].Accept(e)
}

func (e *Executor) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	return nil
}

func (e *Executor) VisitCreateExpr(ctx *CreateExprContext) interface{} {
	return nil
}

func (e *Executor) VisitSelfAddExpr(ctx *SelfAddExprContext) interface{} {
	return ctx.Expression().Accept(e)
}

func (e *Executor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return ctx.Primary().Accept(e)
}

func (e *Executor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return nil
}

func (e *Executor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	cond := e.toBool(ctx.Expression(0).Accept(e))
	if cond {
		return ctx.Expression(1).Accept(e)
	}
	return ctx.Expression(2).Accept(e)
}

func (e *Executor) VisitPrimary(ctx *PrimaryContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(e)
	}
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(e)
	}
	if ctx.Identifier() != nil {
		variable := e.lookupVar(ctx.Identifier().GetText())
		if variable == nil {
			return nil
		}
		return variable.Value.Interface()
	}
	return nil
}

func (e *Executor) VisitLiteral(ctx *LiteralContext) interface{} {
	if ctx.IntegerLiteral() != nil {
		return ctx.IntegerLiteral().Accept(e)
	}
	if ctx.FloatingPointLiteral() != nil {
		value, _ := strconv.ParseFloat(ctx.FloatingPointLiteral().GetText(), 64)
		return value
	}
	if ctx.CharacterLiteral() != nil {
		value, _ := strconv.Unquote(ctx.CharacterLiteral().GetText())
		return value
	}
	if ctx.StringLiteral() != nil {
		value, _ := strconv.Unquote(ctx.StringLiteral().GetText())
		return value
	}
	if ctx.BooleanLiteral() != nil {
		return ctx.BooleanLiteral().GetText() == "true"
	}
	return nil
}

func (e *Executor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	value, _ := strconv.ParseInt(ctx.GetText(), 0, 64)
	return value
}

func (e *Executor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	values := make([]interface{}, 0, len(ctx.AllExpression()))
	for _, expr := range ctx.AllExpression() {
		values = append(values, expr.Accept(e))
	}
	return values
}

func (e *Executor) VisitCreator(ctx *CreatorContext) interface{} {
	return nil
}

func (e *Executor) VisitMapCreator(ctx *MapCreatorContext) interface{} {
	return nil
}

func (e *Executor) VisitArrayCreator(ctx *ArrayCreatorContext) interface{} {
	return nil
}

func (e *Executor) VisitCreatorName(ctx *CreatorNameContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitConnectorCreator(ctx *ConnectorCreatorContext) interface{} {
	return nil
}

func (e *Executor) VisitPrimitiveCreator(ctx *PrimitiveCreatorContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(e)
	}
	return nil
}

func (e *Executor) VisitDynamicCreator(ctx *DynamicCreatorContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(e)
	}
	return nil
}

func (e *Executor) typeFromText(text string) (VariableType, reflect.Type) {
	base := strings.TrimRight(text, "[]")
	switch base {
	case "int", "uint":
		return VarTypeInt, reflect.TypeOf(int64(0))
	case "float":
		return VarTypeFloat, reflect.TypeOf(float64(0))
	case "bool":
		return VarTypeBool, reflect.TypeOf(false)
	case "char":
		return VarTypeChar, reflect.TypeOf("")
	case "string":
		return VarTypeString, reflect.TypeOf("")
	case "dynamic", "error":
		return VarTypeDynamic, reflect.TypeOf((*interface{})(nil)).Elem()
	default:
		return VarTypeDynamic, reflect.TypeOf((*interface{})(nil)).Elem()
	}
}

func (e *Executor) defaultValue(t VariableType) interface{} {
	switch t {
	case VarTypeInt:
		return int64(0)
	case VarTypeFloat:
		return float64(0)
	case VarTypeBool:
		return false
	case VarTypeChar, VarTypeString:
		return ""
	default:
		return nil
	}
}

func (e *Executor) coerceValue(value interface{}, t VariableType) interface{} {
	if value == nil {
		return e.defaultValue(t)
	}
	switch t {
	case VarTypeInt:
		if v, ok := e.toInt64(value); ok {
			return v
		}
	case VarTypeFloat:
		if v, ok := e.toFloat64(value); ok {
			return v
		}
	case VarTypeBool:
		if v, ok := value.(bool); ok {
			return v
		}
	case VarTypeChar, VarTypeString:
		if v, ok := value.(string); ok {
			return v
		}
	}
	return value
}

func (e *Executor) addVar(v *Variable) {
	if e.scopeStack != nil && len(e.scopeStack.stack) > 0 {
		e.scopeStack.stack[len(e.scopeStack.stack)-1].AddVar(v)
		return
	}
	e.varMap[v.Name] = v
}

func (e *Executor) lookupVar(name string) *Variable {
	if e.scopeStack != nil {
		for i := len(e.scopeStack.stack) - 1; i >= 0; i-- {
			if v := e.scopeStack.stack[i].GetVar(name); v != nil {
				return v
			}
		}
	}
	if v, ok := e.varMap[name]; ok {
		return v
	}
	return nil
}

func (e *Executor) toInt64(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int64:
		return v, true
	case uint:
		return int64(v), true
	case uint64:
		return int64(v), true
	case float64:
		return int64(v), true
	}
	return 0, false
}

func (e *Executor) toFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case uint:
		return float64(v), true
	case uint64:
		return float64(v), true
	case float64:
		return v, true
	}
	return 0, false
}

func (e *Executor) toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case int:
		return v != 0
	case float64:
		return v != 0
	}
	return false
}

func (e *Executor) add(left, right interface{}) interface{} {
	if ls, ok := left.(string); ok {
		if rs, ok := right.(string); ok {
			return ls + rs
		}
	}
	if lf, ok := e.toFloat64(left); ok {
		if rf, ok := e.toFloat64(right); ok {
			if _, lok := left.(float64); lok {
				return lf + rf
			}
			if _, rok := right.(float64); rok {
				return lf + rf
			}
			if li, ok := e.toInt64(left); ok {
				if ri, ok := e.toInt64(right); ok {
					return li + ri
				}
			}
			return lf + rf
		}
	}
	return nil
}

func (e *Executor) sub(left, right interface{}) interface{} {
	if _, ok := left.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf - rf
	}
	if _, ok := right.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf - rf
	}
	li, _ := e.toInt64(left)
	ri, _ := e.toInt64(right)
	return li - ri
}

func (e *Executor) mul(left, right interface{}) interface{} {
	if _, ok := left.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf * rf
	}
	if _, ok := right.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf * rf
	}
	li, _ := e.toInt64(left)
	ri, _ := e.toInt64(right)
	return li * ri
}

func (e *Executor) div(left, right interface{}) interface{} {
	if _, ok := left.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf / rf
	}
	if _, ok := right.(float64); ok {
		lf, _ := e.toFloat64(left)
		rf, _ := e.toFloat64(right)
		return lf / rf
	}
	li, _ := e.toInt64(left)
	ri, _ := e.toInt64(right)
	return li / ri
}

func (e *Executor) mod(left, right interface{}) interface{} {
	li, _ := e.toInt64(left)
	ri, _ := e.toInt64(right)
	return li % ri
}

func (e *Executor) negate(value interface{}) interface{} {
	if _, ok := value.(float64); ok {
		f, _ := e.toFloat64(value)
		return -f
	}
	i, _ := e.toInt64(value)
	return -i
}

func (e *Executor) compare(left, right interface{}, op string) bool {
	if ls, ok := left.(string); ok {
		if rs, ok := right.(string); ok {
			switch op {
			case "==":
				return ls == rs
			case "!=":
				return ls != rs
			case "<":
				return ls < rs
			case "<=":
				return ls <= rs
			case ">":
				return ls > rs
			case ">=":
				return ls >= rs
			}
		}
	}
	if lf, ok := e.toFloat64(left); ok {
		if rf, ok := e.toFloat64(right); ok {
			switch op {
			case "==":
				return lf == rf
			case "!=":
				return lf != rf
			case "<":
				return lf < rf
			case "<=":
				return lf <= rf
			case ">":
				return lf > rf
			case ">=":
				return lf >= rf
			}
		}
	}
	if op == "==" {
		return left == right
	}
	if op == "!=" {
		return left != right
	}
	return false
}
