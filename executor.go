package goScript

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Executor struct {
	// Control flags for statement-level flow management.
	breakFlag    bool
	continueFlag bool
	returnFlag   bool
	killFlag     bool
	returnValue  interface{}

	// Runtime state for scopes, functions, and host bindings.
	scopeStack *ScopeStack
	funcStack  *FuncStack
	funcMap    map[string]*FuncStack
	varMap     map[string]*Variable
	GoVarMap   map[string]reflect.Value
	GoTypeMap  map[string]reflect.Type
}

// NewExecutor builds a fresh runtime with a root scope.
func NewExecutor() *Executor {
	executor := &Executor{
		scopeStack: NewScopeStack(),
		funcStack:  NewFuncStack(),
		funcMap:    map[string]*FuncStack{},
		varMap:     map[string]*Variable{},
		GoVarMap:   map[string]reflect.Value{},
		GoTypeMap:  map[string]reflect.Type{},
	}
	executor.scopeStack.Push(NewScope(CommonScopeType))
	return executor
}

// RegisterFunc binds a Go function into the script runtime.
func (e *Executor) RegisterFunc(name string, fn interface{}) {
	e.GoVarMap[name] = reflect.ValueOf(fn)
}

// RegisterConnector binds a Go type prototype for connector<...> creation.
func (e *Executor) RegisterConnector(name string, prototype interface{}) {
	if prototype == nil {
		return
	}
	e.GoTypeMap[name] = reflect.TypeOf(prototype)
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

// VisitChildren stops traversal early when control-flow flags are set.
func (e *Executor) VisitChildren(node antlr.RuleNode) interface{} {
	var result interface{}
	for i := 0; i < node.GetChildCount(); i++ {
		if e.returnFlag || e.breakFlag || e.continueFlag || e.killFlag {
			break
		}
		child := node.GetChild(i)
		if childTree, ok := child.(antlr.ParseTree); ok {
			result = childTree.Accept(e)
		}
		if e.returnFlag || e.breakFlag || e.continueFlag || e.killFlag {
			break
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
	// Top-level: register variables and functions.
	for _, v := range ctx.AllVariableDeclaration() {
		v.Accept(e)
	}
	for _, f := range ctx.AllFunctionDeclaration() {
		f.Accept(e)
	}
	return nil
}

// VisitFunctionDeclaration registers a script-level function definition.
func (e *Executor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	name := ctx.Identifier().GetText()
	var params []*Variable
	if ctx.FormalParameters() != nil {
		for _, p := range ctx.FormalParameters().(*FormalParametersContext).AllFormalParameterDecl() {
			if v, ok := p.Accept(e).(*Variable); ok {
				params = append(params, v)
			}
		}
	}
	var results []*Variable
	if ctx.ReturnType() != nil {
		rt := ctx.ReturnType().(*ReturnTypeContext)
		for _, t := range rt.AllType_() {
			varType, typeGo := e.typeFromText(t.GetText())
			results = append(results, NewVariable("", varType, typeGo, reflect.Zero(typeGo)))
		}
	}
	fn := NewFunction(name, params, results, NewScope(FuncScopeType), ctx.Block().(*BlockContext))
	stack := e.funcMap[name]
	if stack == nil {
		stack = NewFuncStack()
		e.funcMap[name] = stack
	}
	stack.Push(fn)
	return nil
}

func (e *Executor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	params := make([]*Variable, 0, len(ctx.AllFormalParameterDecl()))
	for _, p := range ctx.AllFormalParameterDecl() {
		if v, ok := p.Accept(e).(*Variable); ok {
			params = append(params, v)
		}
	}
	return params
}

func (e *Executor) VisitFormalParameterDecl(ctx *FormalParameterDeclContext) interface{} {
	varType, typeGo := e.typeFromText(ctx.Type_().GetText())
	return NewVariable(ctx.Identifier().GetText(), varType, typeGo, reflect.Zero(typeGo))
}

func (e *Executor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	types := make([]string, 0, len(ctx.AllType_()))
	for _, t := range ctx.AllType_() {
		types = append(types, t.GetText())
	}
	return types
}

func (e *Executor) VisitBlock(ctx *BlockContext) interface{} {
	// Each block introduces a new scope.
	e.PushScope(NewScope(CommonScopeType))
	defer e.PopScope()
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	// Declare typed variables with optional initializer.
	varType, typeGo := e.typeFromText(ctx.Type_().GetText())
	for _, decl := range ctx.VariableDeclarators().AllVariableDeclarator() {
		declCtx := decl.(*VariableDeclaratorContext)
		name := declCtx.Identifier().GetText()
		value := e.defaultValue(varType)
		if declCtx.VariableInitializer() != nil {
			value = declCtx.VariableInitializer().Accept(e)
		}
		actualType := varType
		actualGo := typeGo
		switch value.(type) {
		case []interface{}:
			actualType, actualGo = VarTypeArray, reflect.TypeOf([]interface{}{})
		case map[interface{}]interface{}:
			actualType, actualGo = VarTypeMap, reflect.TypeOf(map[interface{}]interface{}(nil))
		}
		coerced := e.coerceValue(value, actualType)
		var valueRef reflect.Value
		if coerced == nil {
			valueRef = reflect.Zero(actualGo)
		} else {
			valueRef = reflect.ValueOf(coerced)
			if actualGo != nil && valueRef.IsValid() && valueRef.Type() != actualGo && valueRef.Type().ConvertibleTo(actualGo) {
				valueRef = valueRef.Convert(actualGo)
			}
		}
		e.addVar(NewVariable(name, actualType, actualGo, valueRef))
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
	if ctx.ArrayInitializer() != nil {
		return ctx.ArrayInitializer().Accept(e)
	}
	if ctx.MapInitializer() != nil {
		return ctx.MapInitializer().Accept(e)
	}
	return nil
}

func (e *Executor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	items := make([]interface{}, 0, len(ctx.AllVariableInitializer()))
	for _, vi := range ctx.AllVariableInitializer() {
		items = append(items, vi.Accept(e))
	}
	return items
}

func (e *Executor) VisitMapInitializer(ctx *MapInitializerContext) interface{} {
	m := map[interface{}]interface{}{}
	for _, me := range ctx.AllMapEntry() {
		key := me.Expression().Accept(e)
		val := me.VariableInitializer().Accept(e)
		m[key] = val
	}
	return m
}

func (e *Executor) VisitMapEntry(ctx *MapEntryContext) interface{} {
	key := ctx.Expression().Accept(e)
	val := ctx.VariableInitializer().Accept(e)
	return []interface{}{key, val}
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
	// for(init; cond; update) with per-iteration scope and flags.
	ctrl := ctx.ForControl()
	if ctrl != nil {
		if ctrl.ForInit() != nil {
			ctrl.ForInit().Accept(e)
		}
	}
	for {
		if ctrl != nil && ctrl.Expression() != nil {
			if !e.toBool(ctrl.Expression().Accept(e)) {
				break
			}
		}
		e.PushScope(NewScope(ForScopeType))
		e.breakFlag = false
		e.continueFlag = false
		ctx.Statement().Accept(e)
		e.PopScope()
		if e.returnFlag || e.killFlag || e.breakFlag {
			break
		}
		if e.continueFlag {
			e.continueFlag = false
		}
		if ctrl != nil && ctrl.ForUpdate() != nil {
			ctrl.ForUpdate().Accept(e)
		}
	}
	return nil
}

func (e *Executor) VisitForControl(ctx *ForControlContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitForInit(ctx *ForInitContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *Executor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	// Set return flag and capture value for the current function.
	if ctx.Expression() != nil {
		e.returnValue = ctx.Expression().Accept(e)
		e.returnFlag = true
		return e.returnValue
	}
	e.returnValue = nil
	e.returnFlag = true
	return nil
}

func (e *Executor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	e.breakFlag = true
	return true
}

func (e *Executor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	e.continueFlag = true
	return true
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
	ids := ctx.IdentifierList().(*IdentifierListContext).AllIdentifier()
	rhs := ctx.AllExpression()
	values := make([]interface{}, 0, len(rhs))
	for _, r := range rhs {
		values = append(values, r.Accept(e))
	}
	// Expand a single multi-return value into multiple declarations.
	if len(ids) > 1 {
		values = e.expandAssignValues(len(ids), values)
	}
	for i, id := range ids {
		name := id.GetText()
		val := interface{}(nil)
		if i < len(values) {
			val = values[i]
		} else if len(values) > 0 {
			val = values[len(values)-1]
		}
		e.addVar(NewVariable(name, VarTypeDynamic, reflect.TypeOf((*interface{})(nil)).Elem(), e.valueFromInterface(val)))
	}
	if len(values) > 0 {
		return values[len(values)-1]
	}
	return nil
}

func (e *Executor) valueFromInterface(val interface{}) reflect.Value {
	if val == nil {
		return reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
	}
	return reflect.ValueOf(val)
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
	container := ctx.Expression(0).Accept(e)
	index := ctx.Expression(1).Accept(e)
	switch c := container.(type) {
	case []interface{}:
		i, _ := e.toInt64(index)
		if int(i) >= 0 && int(i) < len(c) {
			return c[i]
		}
		return nil
	case map[interface{}]interface{}:
		return c[index]
	case map[string]interface{}:
		if s, ok := index.(string); ok {
			return c[s]
		}
		return nil
	default:
		return nil
	}
}

func (e *Executor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	lhs := ctx.AllLvalue()
	rhs := ctx.AllExpression()
	values := make([]interface{}, 0, len(rhs))
	for _, r := range rhs {
		values = append(values, r.Accept(e))
	}
	// Expand a single multi-return value into multiple assignments.
	if len(lhs) > 1 {
		values = e.expandAssignValues(len(lhs), values)
	}
	write := func(l ILvalueContext, val interface{}) {
		name, steps, ok := e.lvalueToSteps(l)
		if !ok {
			return
		}
		v := e.lookupVar(name)
		if len(steps) == 0 {
			if v == nil {
				e.addVar(NewVariable(name, VarTypeDynamic, reflect.TypeOf((*interface{})(nil)).Elem(), e.valueFromInterface(val)))
			} else {
				v.Value = e.valueFromInterface(val)
			}
			return
		}
		if v == nil {
			return
		}
		updated, ok := e.setLvaluePath(v.Value.Interface(), steps, val)
		if ok {
			v.Value = e.valueFromInterface(updated)
		}
	}
	for i := 0; i < len(lhs); i++ {
		var val interface{}
		if i < len(values) {
			val = values[i]
		} else if len(values) > 0 {
			val = values[len(values)-1]
		}
		write(lhs[i], val)
	}
	if len(values) > 0 {
		return values[len(values)-1]
	}
	return nil
}

type lvalueStep struct {
	kind string
	key  interface{}
}

func (e *Executor) expandAssignValues(lhsCount int, values []interface{}) []interface{} {
	if lhsCount <= 1 || len(values) != 1 {
		return values
	}
	if items, ok := values[0].([]interface{}); ok {
		return items
	}
	return values
}

func (e *Executor) lvalueToSteps(l ILvalueContext) (string, []lvalueStep, bool) {
	if l == nil {
		return "", nil, false
	}
	if l.Lvalue() == nil && l.Expression() == nil && l.Identifier() != nil {
		return l.Identifier().GetText(), nil, true
	}
	if l.Lvalue() != nil {
		name, steps, ok := e.lvalueToSteps(l.Lvalue())
		if !ok {
			return "", nil, false
		}
		if l.Expression() != nil {
			idx := l.Expression().Accept(e)
			return name, append(steps, lvalueStep{kind: "index", key: idx}), true
		}
		if l.Identifier() != nil {
			return name, append(steps, lvalueStep{kind: "field", key: l.Identifier().GetText()}), true
		}
	}
	return "", nil, false
}

func (e *Executor) setLvaluePath(container interface{}, steps []lvalueStep, val interface{}) (interface{}, bool) {
	if len(steps) == 0 {
		return val, true
	}
	step := steps[0]
	last := len(steps) == 1
	switch c := container.(type) {
	case []interface{}:
		if step.kind != "index" {
			return container, false
		}
		i, ok := e.toInt64(step.key)
		if !ok {
			return container, false
		}
		if int(i) < 0 || int(i) >= len(c) {
			return container, false
		}
		if last {
			c[int(i)] = val
			return c, true
		}
		updated, ok := e.setLvaluePath(c[int(i)], steps[1:], val)
		if ok {
			c[int(i)] = updated
			return c, true
		}
		return c, false
	case map[interface{}]interface{}:
		key := step.key
		if last {
			c[key] = val
			return c, true
		}
		child, ok := c[key]
		if !ok {
			return c, false
		}
		updated, ok := e.setLvaluePath(child, steps[1:], val)
		if ok {
			c[key] = updated
			return c, true
		}
		return c, false
	case map[string]interface{}:
		key, ok := step.key.(string)
		if !ok {
			return c, false
		}
		if last {
			c[key] = val
			return c, true
		}
		child, ok := c[key]
		if !ok {
			return c, false
		}
		updated, ok := e.setLvaluePath(child, steps[1:], val)
		if ok {
			c[key] = updated
			return c, true
		}
		return c, false
	}
	rv := reflect.ValueOf(container)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return container, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return container, false
	}
	if step.kind != "field" {
		return container, false
	}
	field, ok := step.key.(string)
	if !ok {
		return container, false
	}
	fv := rv.FieldByName(field)
	if !fv.IsValid() {
		return container, false
	}
	if last {
		if !fv.CanSet() {
			return container, false
		}
		valRv := reflect.ValueOf(val)
		if !valRv.IsValid() {
			fv.Set(reflect.Zero(fv.Type()))
			return container, true
		}
		if valRv.Type().AssignableTo(fv.Type()) {
			fv.Set(valRv)
			return container, true
		}
		if valRv.Type().ConvertibleTo(fv.Type()) {
			fv.Set(valRv.Convert(fv.Type()))
			return container, true
		}
		return container, false
	}
	child := fv.Interface()
	updated, ok := e.setLvaluePath(child, steps[1:], val)
	if !ok {
		return container, false
	}
	if !fv.CanSet() {
		return container, false
	}
	uv := reflect.ValueOf(updated)
	if !uv.IsValid() {
		fv.Set(reflect.Zero(fv.Type()))
		return container, true
	}
	if uv.Type().AssignableTo(fv.Type()) {
		fv.Set(uv)
		return container, true
	}
	if uv.Type().ConvertibleTo(fv.Type()) {
		fv.Set(uv.Convert(fv.Type()))
		return container, true
	}
	return container, false
}

func (e *Executor) VisitSelectorExpr(ctx *SelectorExprContext) interface{} {
	base := ctx.Expression().Accept(e)
	field := ctx.Identifier().GetText()
	switch c := base.(type) {
	case map[interface{}]interface{}:
		return c[field]
	case map[string]interface{}:
		return c[field]
	}
	v := reflect.ValueOf(base)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		f := v.FieldByName(field)
		if f.IsValid() && f.CanInterface() {
			return f.Interface()
		}
	}
	return nil
}

func (e *Executor) VisitCreateExpr(ctx *CreateExprContext) interface{} {
	if ctx.Creator() != nil {
		return ctx.Creator().Accept(e)
	}
	return nil
}

func (e *Executor) VisitSelfAddExpr(ctx *SelfAddExprContext) interface{} {
	// Post ++/-- on assignable expressions.
	expr := ctx.Expression()
	op := "++"
	if strings.HasSuffix(ctx.GetText(), "--") {
		op = "--"
	}
	getter, setter, ok := e.resolveExprLValue(expr)
	if !ok {
		return expr.Accept(e)
	}
	cur := getter()
	if cur == nil {
		return nil
	}
	delta := int64(1)
	if op == "--" {
		delta = -1
	}
	if _, isFloat := cur.(float64); isFloat {
		f, _ := e.toFloat64(cur)
		newVal := f + float64(delta)
		setter(newVal)
		return newVal
	}
	if i, ok := e.toInt64(cur); ok {
		newVal := i + delta
		setter(newVal)
		return newVal
	}
	return cur
}

// resolveExprLValue returns getter/setter for assignable expression forms.
func (e *Executor) resolveExprLValue(expr IExpressionContext) (func() interface{}, func(interface{}), bool) {
	switch t := expr.(type) {
	case *PrimaryExprContext:
		if t.Primary() != nil && t.Primary().Identifier() != nil {
			name := t.Primary().Identifier().GetText()
			v := e.lookupVar(name)
			if v == nil {
				return nil, nil, false
			}
			return func() interface{} {
					if v.Value.IsValid() {
						return v.Value.Interface()
					}
					return nil
				}, func(val interface{}) {
					v.Value = e.valueFromInterface(val)
				}, true
		}
	case *IndexExprContext:
		base := t.Expression(0).Accept(e)
		index := t.Expression(1).Accept(e)
		getter := func() interface{} {
			switch c := base.(type) {
			case []interface{}:
				i, _ := e.toInt64(index)
				if int(i) >= 0 && int(i) < len(c) {
					return c[int(i)]
				}
				return nil
			case map[interface{}]interface{}:
				return c[index]
			case map[string]interface{}:
				if s, ok := index.(string); ok {
					return c[s]
				}
				return nil
			default:
				return nil
			}
		}
		setter := func(val interface{}) {
			switch c := base.(type) {
			case []interface{}:
				i, _ := e.toInt64(index)
				if int(i) >= 0 && int(i) < len(c) {
					c[int(i)] = val
				}
			case map[interface{}]interface{}:
				c[index] = val
			case map[string]interface{}:
				if s, ok := index.(string); ok {
					c[s] = val
				}
			}
		}
		return getter, setter, true
	case *SelectorExprContext:
		base := t.Expression().Accept(e)
		field := t.Identifier().GetText()
		getter := func() interface{} {
			switch c := base.(type) {
			case map[interface{}]interface{}:
				return c[field]
			case map[string]interface{}:
				return c[field]
			}
			v := reflect.ValueOf(base)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				f := v.FieldByName(field)
				if f.IsValid() && f.CanInterface() {
					return f.Interface()
				}
			}
			return nil
		}
		setter := func(val interface{}) {
			switch c := base.(type) {
			case map[interface{}]interface{}:
				c[field] = val
				return
			case map[string]interface{}:
				c[field] = val
				return
			}
			v := reflect.ValueOf(base)
			if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
				f := v.Elem().FieldByName(field)
				if f.IsValid() && f.CanSet() {
					fv := reflect.ValueOf(val)
					if fv.IsValid() && fv.Type().AssignableTo(f.Type()) {
						f.Set(fv)
					} else if fv.IsValid() && fv.Type().ConvertibleTo(f.Type()) {
						f.Set(fv.Convert(f.Type()))
					}
				}
			}
		}
		return getter, setter, true
	}
	return nil, nil, false
}

func (e *Executor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return ctx.Primary().Accept(e)
}

func (e *Executor) VisitCallExpr(ctx *CallExprContext) interface{} {
	// Call order: selector-method, script function, Go binding.
	if sel, ok := ctx.Expression().(*SelectorExprContext); ok {
		base := sel.Expression().Accept(e)
		field := sel.Identifier().GetText()
		if caller := NewGoCaller(base); caller != nil {
			if out, ok := caller.CallFunc(field, e.evalArgs(ctx.ExpressionList())); ok {
				return out
			}
		}
		bv := reflect.ValueOf(base)
		if bv.IsValid() {
			method := bv.MethodByName(field)
			if !method.IsValid() && bv.Kind() != reflect.Ptr && bv.CanAddr() {
				method = bv.Addr().MethodByName(field)
			}
			if method.IsValid() {
				return e.callReflectFunc(method, ctx.ExpressionList())
			}
		}
		if m, ok := base.(map[string]interface{}); ok {
			if fv, ok := m[field]; ok {
				return e.callReflectFunc(reflect.ValueOf(fv), ctx.ExpressionList())
			}
		}
		if m, ok := base.(map[interface{}]interface{}); ok {
			if fv, ok := m[field]; ok {
				return e.callReflectFunc(reflect.ValueOf(fv), ctx.ExpressionList())
			}
		}
	}
	callee := ctx.Expression().GetText()
	if fn := e.getScriptFunction(callee); fn != nil {
		return e.callScriptFunction(fn, ctx.ExpressionList())
	}
	if fn, ok := e.GoVarMap[callee]; ok {
		return e.callReflectFunc(fn, ctx.ExpressionList())
	}
	return nil
}

func (e *Executor) getScriptFunction(name string) *Function {
	if stack, ok := e.funcMap[name]; ok && stack != nil && !stack.IsEmpty() {
		return stack.Top()
	}
	return nil
}

func (e *Executor) callScriptFunction(fn *Function, list IExpressionListContext) interface{} {
	// Execute script function in its own function scope.
	args := e.evalArgs(list)
	callFn := NewFunction(fn.Name, e.cloneVars(fn.ParametersList), e.cloneVars(fn.ResultsList), NewScope(FuncScopeType), fn.Block)
	for i, param := range callFn.ParametersList {
		var val interface{}
		if i < len(args) {
			val = args[i]
		}
		coerced := e.coerceValue(val, param.Type)
		if coerced == nil {
			param.Value = reflect.Zero(param.TypeGo)
		} else {
			rv := reflect.ValueOf(coerced)
			if param.TypeGo != nil && rv.IsValid() && rv.Type() != param.TypeGo && rv.Type().ConvertibleTo(param.TypeGo) {
				rv = rv.Convert(param.TypeGo)
			}
			param.Value = rv
		}
		callFn.Scope.AddVar(param)
	}
	prevReturnFlag := e.returnFlag
	prevReturnValue := e.returnValue
	prevBreakFlag := e.breakFlag
	prevContinueFlag := e.continueFlag
	prevKillFlag := e.killFlag
	e.returnFlag = false
	e.returnValue = nil
	e.breakFlag = false
	e.continueFlag = false
	e.killFlag = false
	e.PushFunc(callFn)
	callFn.Block.Accept(e)
	e.PopFunc()
	ret := e.returnValue
	e.returnFlag = prevReturnFlag
	e.returnValue = prevReturnValue
	e.breakFlag = prevBreakFlag
	e.continueFlag = prevContinueFlag
	e.killFlag = prevKillFlag
	if len(callFn.ResultsList) == 0 {
		return nil
	}
	if len(callFn.ResultsList) == 1 {
		return e.coerceValue(ret, callFn.ResultsList[0].Type)
	}
	if items, ok := ret.([]interface{}); ok {
		out := make([]interface{}, 0, len(callFn.ResultsList))
		for i, res := range callFn.ResultsList {
			var v interface{}
			if i < len(items) {
				v = e.coerceValue(items[i], res.Type)
			}
			out = append(out, v)
		}
		return out
	}
	return ret
}

func (e *Executor) cloneVars(vars []*Variable) []*Variable {
	out := make([]*Variable, 0, len(vars))
	for _, v := range vars {
		out = append(out, NewVariable(v.Name, v.Type, v.TypeGo, reflect.Zero(v.TypeGo)))
	}
	return out
}

func (e *Executor) evalArgs(list IExpressionListContext) []interface{} {
	if list == nil {
		return nil
	}
	exprs := list.(*ExpressionListContext).AllExpression()
	args := make([]interface{}, 0, len(exprs))
	for _, expr := range exprs {
		args = append(args, expr.Accept(e))
	}
	return args
}

func (e *Executor) callReflectFunc(fn reflect.Value, list IExpressionListContext) interface{} {
	// Execute Go function via reflection with basic coercion.
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil
	}
	var args []reflect.Value
	if list != nil {
		exprs := list.(*ExpressionListContext).AllExpression()
		args = make([]reflect.Value, 0, len(exprs))
		for i, expr := range exprs {
			val := expr.Accept(e)
			rv := reflect.ValueOf(val)
			if i < fn.Type().NumIn() {
				param := fn.Type().In(i)
				if !rv.IsValid() {
					rv = reflect.Zero(param)
				} else if rv.Type().AssignableTo(param) {
				} else if rv.Type().ConvertibleTo(param) {
					rv = rv.Convert(param)
				} else if param.Kind() == reflect.Interface {
				} else {
					rv = reflect.Zero(param)
				}
			}
			args = append(args, rv)
		}
	}
	results := fn.Call(args)
	if len(results) == 0 {
		return nil
	}
	if len(results) == 1 {
		return results[0].Interface()
	}
	out := make([]interface{}, 0, len(results))
	for _, r := range results {
		out = append(out, r.Interface())
	}
	return out
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
		if variable == nil || !variable.Value.IsValid() {
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

func (e *Executor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	ids := ctx.AllIdentifier()
	out := make([]string, 0, len(ids))
	for _, id := range ids {
		out = append(out, id.GetText())
	}
	return out
}

func (e *Executor) VisitLvalue(ctx *LvalueContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitCreator(ctx *CreatorContext) interface{} { return e.VisitChildren(ctx) }

func (e *Executor) VisitMapCreator(ctx *MapCreatorContext) interface{} {
	return ctx.MapInitializer().Accept(e)
}

func (e *Executor) VisitArrayCreator(ctx *ArrayCreatorContext) interface{} {
	return ctx.ArrayInitializer().Accept(e)
}

func (e *Executor) VisitCreatorName(ctx *CreatorNameContext) interface{} {
	return ctx.GetText()
}

func (e *Executor) VisitConnectorCreator(ctx *ConnectorCreatorContext) interface{} {
	// Instantiate a connector by registered Go type prototype.
	if ctx.ConnectorType() == nil {
		return nil
	}
	name := ctx.ConnectorType().(*ConnectorTypeContext).Identifier().GetText()
	t, ok := e.GoTypeMap[name]
	if !ok || t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()).Interface()
	}
	return reflect.New(t).Interface()
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
	// Map grammar type text to internal variable type and Go type.
	if strings.Contains(text, "map<") {
		return VarTypeMap, reflect.TypeOf(map[interface{}]interface{}(nil))
	}
	if strings.Contains(text, "[]") {
		return VarTypeArray, reflect.TypeOf([]interface{}{})
	}
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
	case VarTypeArray:
		return []interface{}{}
	case VarTypeMap:
		return map[interface{}]interface{}{}
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
	// Add variable into the current scope, or fallback to global map.
	if e.scopeStack != nil && len(e.scopeStack.stack) > 0 {
		e.scopeStack.stack[len(e.scopeStack.stack)-1].AddVar(v)
		return
	}
	e.varMap[v.Name] = v
}

func (e *Executor) lookupVar(name string) *Variable {
	// Resolve variable by walking scopes from inner to outer.
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
