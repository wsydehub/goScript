package goScript

import (
	"reflect"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type SampleConnector struct {
	Count int64
}

func (c *SampleConnector) Inc() {
	c.Count++
}

func (c *SampleConnector) Add(v int64) int64 {
	c.Count += v
	return c.Count
}

func runCompilationUnit(executor *Executor, code string) {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	tree := parser.CompilationUnit()
	tree.Accept(executor)
}

func runBlockStatement(executor *Executor, code string) {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	stmt := parser.BlockStatement().(*BlockStatementContext)
	executor.VisitBlockStatement(stmt)
}

func runVariableDeclaration(executor *Executor, code string) {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	decl := parser.VariableDeclaration().(*VariableDeclarationContext)
	executor.VisitVariableDeclaration(decl)
}

func runStatement(executor *Executor, code string) {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	stmt := parser.Statement()
	stmt.Accept(executor)
}

func toInt64Test(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int64:
		return v
	case float64:
		return int64(v)
	}
	return 0
}

func getVar(executor *Executor, name string) *Variable {
	if v := executor.lookupVar(name); v != nil {
		return v
	}
	if v, ok := executor.varMap[name]; ok {
		return v
	}
	if executor.scopeStack != nil && len(executor.scopeStack.stack) > 0 {
		if v := executor.scopeStack.stack[0].GetVar(name); v != nil {
			return v
		}
		last := executor.scopeStack.stack[len(executor.scopeStack.stack)-1]
		if v := last.GetVar(name); v != nil {
			return v
		}
	}
	return nil
}

func TestBasicArithmeticVar(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "a := 1 + 2;")
	v := getVar(executor, "a")
	if v == nil {
		t.Fatalf("var a not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestArrayAssign(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "arr := new int[] {1,2,3};")
	runStatement(executor, "arr[1] = 5;")
	v := getVar(executor, "arr")
	if v == nil {
		t.Fatalf("var arr not found")
	}
	arr, ok := v.Value.Interface().([]interface{})
	if !ok {
		t.Fatalf("arr type invalid")
	}
	if toInt64Test(arr[1]) != 5 {
		t.Fatalf("expect 5, got %v", arr[1])
	}
}

func TestMapAssign(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "m := new map<int,int> {1:2,3:4};")
	runStatement(executor, "m[3] = 7;")
	v := getVar(executor, "m")
	if v == nil {
		t.Fatalf("var m not found")
	}
	m, ok := v.Value.Interface().(map[interface{}]interface{})
	if !ok {
		t.Fatalf("m type invalid")
	}
	if toInt64Test(m[int64(3)]) != 7 {
		t.Fatalf("expect 7, got %v", m[int64(3)])
	}
}

func TestForBreak(t *testing.T) {
	executor := NewExecutor()
	executor.addVar(NewVariable("i", VarTypeInt, reflect.TypeOf(int64(0)), reflect.ValueOf(int64(0))))
	runStatement(executor, "for(; i < 5; i = i + 1) { if i == 3 { break; } }")
	v := getVar(executor, "i")
	if v == nil {
		t.Fatalf("var i not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestCallExpr(t *testing.T) {
	executor := NewExecutor()
	executor.RegisterFunc("add", func(a int64, b int64) int64 { return a + b })
	runStatement(executor, "x := 0;")
	runStatement(executor, "x = add(1,2);")
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestVariableInitializer(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "int[] arr = {1,2}; map<int,int> m = {1:2,3:4};")
	arrVar := getVar(executor, "arr")
	if arrVar == nil {
		t.Fatalf("var arr not found")
	}
	arr, ok := arrVar.Value.Interface().([]interface{})
	if !ok {
		t.Fatalf("arr type invalid")
	}
	if toInt64Test(arr[0]) != 1 || toInt64Test(arr[1]) != 2 {
		t.Fatalf("arr value invalid")
	}
	mVar := getVar(executor, "m")
	if mVar == nil {
		t.Fatalf("var m not found")
	}
	m, ok := mVar.Value.Interface().(map[interface{}]interface{})
	if !ok {
		t.Fatalf("m type invalid")
	}
	if toInt64Test(m[int64(1)]) != 2 || toInt64Test(m[int64(3)]) != 4 {
		t.Fatalf("map value invalid")
	}
}

func TestContinue(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "i := 0;")
	runStatement(executor, "sum := 0;")
	runStatement(executor, "for(; i < 5; i = i + 1) { if i == 2 { continue; } sum = sum + i; }")
	v := getVar(executor, "sum")
	if v == nil {
		t.Fatalf("var sum not found")
	}
	if toInt64Test(v.Value.Interface()) != 8 {
		t.Fatalf("expect 8, got %v", v.Value.Interface())
	}
}

func TestSelfAddSub(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "x := 2;")
	runStatement(executor, "x--;")
	runStatement(executor, "x++;")
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 2 {
		t.Fatalf("expect 2, got %v", v.Value.Interface())
	}
	runStatement(executor, "arr := new int[] {1,2};")
	runStatement(executor, "arr[0]++;")
	arrVar := getVar(executor, "arr")
	if arrVar == nil {
		t.Fatalf("var arr not found")
	}
	arr, ok := arrVar.Value.Interface().([]interface{})
	if !ok {
		t.Fatalf("arr type invalid")
	}
	if toInt64Test(arr[0]) != 2 {
		t.Fatalf("expect 2, got %v", arr[0])
	}
}

func TestScriptFunction(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "func add(int a, int b) int { return a + b; }")
	runStatement(executor, "y := add(1,2);")
	v := getVar(executor, "y")
	if v == nil {
		t.Fatalf("var y not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestConnector(t *testing.T) {
	executor := NewExecutor()
	executor.RegisterConnector("SampleConnector", &SampleConnector{})
	runStatement(executor, "c := new connector<SampleConnector>();")
	runStatement(executor, "c.Inc();")
	runStatement(executor, "r := c.Add(3);")
	runStatement(executor, "z := c.Count;")
	r := getVar(executor, "r")
	z := getVar(executor, "z")
	if r == nil || z == nil {
		t.Fatalf("var r/z not found")
	}
	if toInt64Test(r.Value.Interface()) != 4 {
		t.Fatalf("expect 4, got %v", r.Value.Interface())
	}
	if toInt64Test(z.Value.Interface()) != 4 {
		t.Fatalf("expect 4, got %v", z.Value.Interface())
	}
}

func TestComplexLValueAssign(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, `m := new map<string,dynamic> {"a": new map<string,dynamic> {"b": new int[] {1,2}}};`)
	runStatement(executor, `m["a"].b[1] = 7;`)
	v := getVar(executor, "m")
	if v == nil {
		t.Fatalf("var m not found")
	}
	root, ok := v.Value.Interface().(map[interface{}]interface{})
	if !ok {
		t.Fatalf("m type invalid")
	}
	a, ok := root["a"].(map[interface{}]interface{})
	if !ok {
		t.Fatalf("nested map invalid")
	}
	b, ok := a["b"].([]interface{})
	if !ok {
		t.Fatalf("nested array invalid")
	}
	if toInt64Test(b[1]) != 7 {
		t.Fatalf("expect 7, got %v", b[1])
	}
}

func TestMultiReturnAssign(t *testing.T) {
	executor := NewExecutor()
	executor.RegisterFunc("pair", func(a int64) (int64, int64) { return a, a + 1 })
	runStatement(executor, "x := 0; y := 0;")
	runStatement(executor, "x, y = pair(1);")
	vx := getVar(executor, "x")
	vy := getVar(executor, "y")
	if vx == nil || vy == nil {
		t.Fatalf("vars x/y not found")
	}
	if toInt64Test(vx.Value.Interface()) != 1 || toInt64Test(vy.Value.Interface()) != 2 {
		t.Fatalf("expect 1,2 got %v,%v", vx.Value.Interface(), vy.Value.Interface())
	}
	runStatement(executor, "a, b := pair(3);")
	va := getVar(executor, "a")
	vb := getVar(executor, "b")
	if va == nil || vb == nil {
		t.Fatalf("vars a/b not found")
	}
	if toInt64Test(va.Value.Interface()) != 3 || toInt64Test(vb.Value.Interface()) != 4 {
		t.Fatalf("expect 3,4 got %v,%v", va.Value.Interface(), vb.Value.Interface())
	}
	runCompilationUnit(executor, "func pair2(int a) (int,int) { return new int[] {a, a + 2}; }")
	runStatement(executor, "m, n := pair2(5);")
	vm := getVar(executor, "m")
	vn := getVar(executor, "n")
	if vm == nil || vn == nil {
		t.Fatalf("vars m/n not found")
	}
	if toInt64Test(vm.Value.Interface()) != 5 || toInt64Test(vn.Value.Interface()) != 7 {
		t.Fatalf("expect 5,7 got %v,%v", vm.Value.Interface(), vn.Value.Interface())
	}
}
