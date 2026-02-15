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

func runBlock(executor *Executor, code string) {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	block := parser.Block().(*BlockContext)
	executor.VisitBlock(block)
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

func parseCompilationHasError(code string) bool {
	input := antlr.NewInputStream(code)
	lexer := NewGoScriptLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewGoScriptParser(tokens)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	listener := &parseErrorListener{}
	parser.AddErrorListener(listener)
	parser.CompilationUnit()
	return listener.count > 0
}

type parseErrorListener struct {
	count int
}

func (p *parseErrorListener) SyntaxError(antlr.Recognizer, interface{}, int, int, string, antlr.RecognitionException) {
	p.count++
}

func (p *parseErrorListener) ReportAmbiguity(antlr.Parser, *antlr.DFA, int, int, bool, *antlr.BitSet, *antlr.ATNConfigSet) {
}

func (p *parseErrorListener) ReportAttemptingFullContext(antlr.Parser, *antlr.DFA, int, int, *antlr.BitSet, *antlr.ATNConfigSet) {
}

func (p *parseErrorListener) ReportContextSensitivity(antlr.Parser, *antlr.DFA, int, int, int, *antlr.ATNConfigSet) {
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

func TestTypeParsingAndCreators(t *testing.T) {
	executor := NewExecutor()
	executor.RegisterConnector("SampleConnector", &SampleConnector{})
	runCompilationUnit(executor, "map<int,int>[] ms = {{1:2},{3:4}};")
	ms := getVar(executor, "ms")
	if ms == nil {
		t.Fatalf("var ms not found")
	}
	arr, ok := ms.Value.Interface().([]interface{})
	if !ok || len(arr) != 2 {
		t.Fatalf("ms type invalid")
	}
	m0, ok := arr[0].(map[interface{}]interface{})
	if !ok || toInt64Test(m0[int64(1)]) != 2 {
		t.Fatalf("ms[0] invalid")
	}
	runStatement(executor, "x := new int(1);")
	x := getVar(executor, "x")
	if x == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(x.Value.Interface()) != 1 {
		t.Fatalf("expect 1, got %v", x.Value.Interface())
	}
	runStatement(executor, "y := new float(1);")
	y := getVar(executor, "y")
	if y == nil {
		t.Fatalf("var y not found")
	}
	if _, ok := y.Value.Interface().(float64); !ok {
		t.Fatalf("expect float64, got %T", y.Value.Interface())
	}
	runBlockStatement(executor, "connector<SampleConnector> c = new connector<SampleConnector>();")
	runStatement(executor, "c.Inc();")
	runStatement(executor, "z := c.Count;")
	z := getVar(executor, "z")
	if z == nil {
		t.Fatalf("var z not found")
	}
	if toInt64Test(z.Value.Interface()) != 1 {
		t.Fatalf("expect 1, got %v", z.Value.Interface())
	}
}

func TestIfElse(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "x := 0;")
	runStatement(executor, "if 1 < 2 { x = 3; } else { x = 4; }")
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
	runStatement(executor, "if 2 < 1 { x = 5; } else { x = 6; }")
	if toInt64Test(v.Value.Interface()) != 6 {
		t.Fatalf("expect 6, got %v", v.Value.Interface())
	}
}

func TestForInitUpdate(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "sum := 0;")
	runStatement(executor, "for(int i = 0; i < 3; i = i + 1) { sum = sum + i; }")
	v := getVar(executor, "sum")
	if v == nil {
		t.Fatalf("var sum not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestExpressionOps(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "a := 1 + 2 * 3;")
	runStatement(executor, "b := (1 + 2) * 3;")
	runStatement(executor, "c := 5 % 2;")
	runStatement(executor, "d := 2 > 1;")
	runStatement(executor, "e := 2 <= 1;")
	runStatement(executor, "f := !false;")
	runStatement(executor, "g := true && false;")
	runStatement(executor, "h := false || true;")
	runStatement(executor, "i := 1 < 2 ? 3 : 4;")
	a := getVar(executor, "a")
	b := getVar(executor, "b")
	c := getVar(executor, "c")
	d := getVar(executor, "d")
	e := getVar(executor, "e")
	f := getVar(executor, "f")
	g := getVar(executor, "g")
	h := getVar(executor, "h")
	i := getVar(executor, "i")
	if a == nil || b == nil || c == nil || d == nil || e == nil || f == nil || g == nil || h == nil || i == nil {
		t.Fatalf("vars not found")
	}
	if toInt64Test(a.Value.Interface()) != 7 || toInt64Test(b.Value.Interface()) != 9 {
		t.Fatalf("arith invalid")
	}
	if toInt64Test(c.Value.Interface()) != 1 {
		t.Fatalf("mod invalid")
	}
	if d.Value.Interface() != true || e.Value.Interface() != false {
		t.Fatalf("compare invalid")
	}
	if f.Value.Interface() != true || g.Value.Interface() != false || h.Value.Interface() != true {
		t.Fatalf("logic invalid")
	}
	if toInt64Test(i.Value.Interface()) != 3 {
		t.Fatalf("ternary invalid")
	}
}

func TestLiterals(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, `a := 0x10;`)
	runStatement(executor, `b := 010;`)
	runStatement(executor, `c := 10;`)
	runStatement(executor, `d := 1.5;`)
	runStatement(executor, `e := 'a';`)
	runStatement(executor, `f := "hi";`)
	runStatement(executor, `g := true;`)
	runStatement(executor, `h := false;`)
	runStatement(executor, `i := null;`)
	a := getVar(executor, "a")
	b := getVar(executor, "b")
	c := getVar(executor, "c")
	d := getVar(executor, "d")
	e := getVar(executor, "e")
	f := getVar(executor, "f")
	g := getVar(executor, "g")
	h := getVar(executor, "h")
	i := getVar(executor, "i")
	if a == nil || b == nil || c == nil || d == nil || e == nil || f == nil || g == nil || h == nil || i == nil {
		t.Fatalf("vars not found")
	}
	if toInt64Test(a.Value.Interface()) != 16 || toInt64Test(b.Value.Interface()) != 8 || toInt64Test(c.Value.Interface()) != 10 {
		t.Fatalf("int literal invalid")
	}
	if _, ok := d.Value.Interface().(float64); !ok {
		t.Fatalf("float literal invalid")
	}
	if e.Value.Interface() != "a" || f.Value.Interface() != "hi" {
		t.Fatalf("string literal invalid")
	}
	if g.Value.Interface() != true || h.Value.Interface() != false {
		t.Fatalf("bool literal invalid")
	}
	if i.Value.Interface() != nil {
		t.Fatalf("null literal invalid")
	}
}

func TestCreateAndAssign(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, "a, b := 1, 2;")
	a := getVar(executor, "a")
	b := getVar(executor, "b")
	if a == nil || b == nil {
		t.Fatalf("vars a/b not found")
	}
	if toInt64Test(a.Value.Interface()) != 1 || toInt64Test(b.Value.Interface()) != 2 {
		t.Fatalf("assign invalid")
	}
}

func TestVariableDeclarationDefaults(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "int a; float b; bool c; char d; string e; dynamic f; map<int,int> g; int[] h;")
	a := getVar(executor, "a")
	b := getVar(executor, "b")
	c := getVar(executor, "c")
	d := getVar(executor, "d")
	e := getVar(executor, "e")
	f := getVar(executor, "f")
	g := getVar(executor, "g")
	h := getVar(executor, "h")
	if a == nil || b == nil || c == nil || d == nil || e == nil || f == nil || g == nil || h == nil {
		t.Fatalf("vars not found")
	}
	if toInt64Test(a.Value.Interface()) != 0 {
		t.Fatalf("expect 0, got %v", a.Value.Interface())
	}
	if _, ok := b.Value.Interface().(float64); !ok {
		t.Fatalf("expect float64, got %T", b.Value.Interface())
	}
	if c.Value.Interface() != false {
		t.Fatalf("expect false, got %v", c.Value.Interface())
	}
	if d.Value.Interface() != "" || e.Value.Interface() != "" {
		t.Fatalf("expect empty string, got %v/%v", d.Value.Interface(), e.Value.Interface())
	}
	if f.Value.Interface() != nil {
		t.Fatalf("expect nil, got %v", f.Value.Interface())
	}
	if _, ok := g.Value.Interface().(map[interface{}]interface{}); !ok {
		t.Fatalf("expect map, got %T", g.Value.Interface())
	}
	if _, ok := h.Value.Interface().([]interface{}); !ok {
		t.Fatalf("expect array, got %T", h.Value.Interface())
	}
}

func TestVariableDeclarationInit(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, `
int a = 1;
float b = 1;
bool c = true;
char d = 'x';
string e = "hi";
dynamic f = 3;
map<int,int> g = {1:2,3:4};
int[] h = {5,6};
`)
	a := getVar(executor, "a")
	b := getVar(executor, "b")
	c := getVar(executor, "c")
	d := getVar(executor, "d")
	e := getVar(executor, "e")
	f := getVar(executor, "f")
	g := getVar(executor, "g")
	h := getVar(executor, "h")
	if a == nil || b == nil || c == nil || d == nil || e == nil || f == nil || g == nil || h == nil {
		t.Fatalf("vars not found")
	}
	if toInt64Test(a.Value.Interface()) != 1 {
		t.Fatalf("expect 1, got %v", a.Value.Interface())
	}
	if _, ok := b.Value.Interface().(float64); !ok {
		t.Fatalf("expect float64, got %T", b.Value.Interface())
	}
	if c.Value.Interface() != true {
		t.Fatalf("expect true, got %v", c.Value.Interface())
	}
	if d.Value.Interface() != "x" || e.Value.Interface() != "hi" {
		t.Fatalf("expect string values, got %v/%v", d.Value.Interface(), e.Value.Interface())
	}
	if toInt64Test(f.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", f.Value.Interface())
	}
	gm, ok := g.Value.Interface().(map[interface{}]interface{})
	if !ok || toInt64Test(gm[int64(1)]) != 2 || toInt64Test(gm[int64(3)]) != 4 {
		t.Fatalf("map init invalid")
	}
	ha, ok := h.Value.Interface().([]interface{})
	if !ok || toInt64Test(ha[0]) != 5 || toInt64Test(ha[1]) != 6 {
		t.Fatalf("array init invalid")
	}
}

func TestMultiDimArrayRuntimeMeta(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "int[][] grid = {{1,2},{3,4}};")
	grid := getVar(executor, "grid")
	if grid == nil {
		t.Fatalf("var grid not found")
	}
	if grid.Type != VarTypeArray {
		t.Fatalf("expect array type, got %v", grid.Type)
	}
	if grid.ArrayDims != 2 {
		t.Fatalf("expect 2 dims, got %d", grid.ArrayDims)
	}
}

func TestSelectorAndIndexRead(t *testing.T) {
	executor := NewExecutor()
	runStatement(executor, `m := new map<string,int> {"a": 1};`)
	runStatement(executor, `arr := new int[] {9,8};`)
	runStatement(executor, `x := m["a"];`)
	runStatement(executor, `y := arr[1];`)
	x := getVar(executor, "x")
	y := getVar(executor, "y")
	if x == nil || y == nil {
		t.Fatalf("vars x/y not found")
	}
	if toInt64Test(x.Value.Interface()) != 1 || toInt64Test(y.Value.Interface()) != 8 {
		t.Fatalf("read invalid")
	}
}

func TestLongScriptScenario(t *testing.T) {
	executor := NewExecutor()
	executor.RegisterFunc("add", func(a int64, b int64) int64 { return a + b })
	executor.RegisterConnector("SampleConnector", &SampleConnector{})
	runCompilationUnit(executor, `
func fib(int n) int {
    if n <= 1 { return n; }
    return fib(n - 1) + fib(n - 2);
}
func pair(int a) (int, int) {
    return new int[] {a, a + 1};
}
int base = 2;
int out = 0;
`)
	runBlockStatement(executor, `
{
    sum := 0;
    arr := new int[] {1,2,3};
    m := new map<string,int> {"a": 1, "b": 2};
    for(int i = 0; i < 3; i = i + 1) {
        sum = sum + arr[i];
    }
    if sum > 5 { m["c"] = sum; } else { m["c"] = 0; }
    x, y := pair(sum);
    z := fib(5);
    k := add(base, y);
    c := new connector<SampleConnector>();
    c.Inc();
    c.Add(3);
    result := (m["c"] + x + y + z + k + c.Count) > 20 ? 1 : 0;
    out = result;
}
`)
	out := getVar(executor, "out")
	if out == nil {
		t.Fatalf("var out not found")
	}
	if toInt64Test(out.Value.Interface()) != 1 {
		t.Fatalf("expect 1, got %v", out.Value.Interface())
	}
}

// TestPathPlanningScript validates a grid shortest path via relaxation in GoScript.
func TestPathPlanningScript(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "int distOut = 0;")
	runBlockStatement(executor, `
{
    rows := 4;
    cols := 4;
    startR := 0;
    startC := 0;
    targetR := 3;
    targetC := 3;
    obstacles := new map<int,int> {11:1, 21:1};
    dist := new int[] {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0};
    inf := 999;
    r0 := 0;
    for(; r0 < rows; r0 = r0 + 1) {
        c0 := 0;
        for(; c0 < cols; c0 = c0 + 1) {
            dist[r0 * cols + c0] = inf;
        }
    }
    dist[startR * cols + startC] = 0;
    dr := new int[] {1, -1, 0, 0};
    dc := new int[] {0, 0, 1, -1};
    step := 0;
    for(; step < rows * cols; step = step + 1) {
        r := 0;
        for(; r < rows; r = r + 1) {
            c := 0;
            for(; c < cols; c = c + 1) {
            key := r * 10 + c;
            idx := r * cols + c;
            if obstacles[key] == 1 { continue; }
            if dist[idx] == inf { continue; }
                i := 0;
                for(; i < 4; i = i + 1) {
                    nr := r + dr[i];
                    nc := c + dc[i];
                    if nr < 0 || nr >= rows || nc < 0 || nc >= cols { continue; }
                    nkey := nr * 10 + nc;
                nidx := nr * cols + nc;
                    if obstacles[nkey] == 1 { continue; }
                if dist[nidx] > dist[idx] + 1 {
                    dist[nidx] = dist[idx] + 1;
                    }
                }
            }
        }
    }
    distOut = dist[targetR * cols + targetC];
}
`)
	out := getVar(executor, "distOut")
	if out == nil {
		t.Fatalf("var distOut not found")
	}
	if toInt64Test(out.Value.Interface()) != 6 {
		t.Fatalf("expect 6, got %v", out.Value.Interface())
	}
}

// TestPathPlanningScriptWithMultiDimArray validates grid shortest path using int[][].
func TestPathPlanningScriptWithMultiDimArray(t *testing.T) {
	executor := NewExecutor()
	runCompilationUnit(executor, "int distOut2 = 0;")
	runBlockStatement(executor, `
{
    rows := 4;
    cols := 4;
    startR := 0;
    startC := 0;
    targetR := 3;
    targetC := 3;
    grid := new int[][] {
        {0,0,0,0},
        {0,1,0,0},
        {0,1,0,0},
        {0,0,0,0}
    };
    dist := new int[][] {
        {0,0,0,0},
        {0,0,0,0},
        {0,0,0,0},
        {0,0,0,0}
    };
    inf := 999;
    r0 := 0;
    for(; r0 < rows; r0 = r0 + 1) {
        c0 := 0;
        for(; c0 < cols; c0 = c0 + 1) {
            dist[r0][c0] = inf;
        }
    }
    dist[startR][startC] = 0;
    dr := new int[] {1, -1, 0, 0};
    dc := new int[] {0, 0, 1, -1};
    step := 0;
    for(; step < rows * cols; step = step + 1) {
        r := 0;
        for(; r < rows; r = r + 1) {
            c := 0;
            for(; c < cols; c = c + 1) {
                if grid[r][c] == 1 { continue; }
                if dist[r][c] == inf { continue; }
                i := 0;
                for(; i < 4; i = i + 1) {
                    nr := r + dr[i];
                    nc := c + dc[i];
                    if nr < 0 || nr >= rows || nc < 0 || nc >= cols { continue; }
                    if grid[nr][nc] == 1 { continue; }
                    if dist[nr][nc] > dist[r][c] + 1 {
                        dist[nr][nc] = dist[r][c] + 1;
                    }
                }
            }
        }
    }
    distOut2 = dist[targetR][targetC];
}
`)
	out := getVar(executor, "distOut2")
	if out == nil {
		t.Fatalf("var distOut2 not found")
	}
	if toInt64Test(out.Value.Interface()) != 6 {
		t.Fatalf("expect 6, got %v", out.Value.Interface())
	}
}

func TestForLoop(t *testing.T) {
	executor := NewExecutor()
	code := `
sum := 0;
loopCount := 10000;
for(int i = 0; i < loopCount; i = i + 1) {
    sum = sum + i;
}
	`
	runCompilationUnit(executor, code)
	v := getVar(executor, "sum")
	if v == nil {
		t.Fatalf("var sum not found")
	}
	if toInt64Test(v.Value.Interface()) != 49995000 {
		t.Fatalf("expect 49995000, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementRestrictions(t *testing.T) {
	if !parseCompilationHasError(`return 1;`) {
		t.Fatalf("expect parse errors for top-level return")
	}
}

func TestTopLevelStatementBlock(t *testing.T) {
	executor := NewExecutor()
	code := `
int sum = 0;
{
    sum = 2;
}
`
	runCompilationUnit(executor, code)
	v := getVar(executor, "sum")
	if v == nil {
		t.Fatalf("var sum not found")
	}
	if toInt64Test(v.Value.Interface()) != 2 {
		t.Fatalf("expect 2, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementIf(t *testing.T) {
	executor := NewExecutor()
	code := `
int x = 0;
if true {
    x = 1;
}
`
	runCompilationUnit(executor, code)
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 1 {
		t.Fatalf("expect 1, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementFor(t *testing.T) {
	executor := NewExecutor()
	code := `
int s = 0;
for(int i = 0; i < 3; i = i + 1) {
    s = s + 1;
}
`
	runCompilationUnit(executor, code)
	v := getVar(executor, "s")
	if v == nil {
		t.Fatalf("var s not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementExpression(t *testing.T) {
	executor := NewExecutor()
	code := `
int x = 1;
x = x + 2;
`
	runCompilationUnit(executor, code)
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementFunctionCall(t *testing.T) {
	executor := NewExecutor()
	code := `
func add(int a, int b) int { return a + b; }
x := add(1, 2);
`
	runCompilationUnit(executor, code)
	v := getVar(executor, "x")
	if v == nil {
		t.Fatalf("var x not found")
	}
	if toInt64Test(v.Value.Interface()) != 3 {
		t.Fatalf("expect 3, got %v", v.Value.Interface())
	}
}

func TestTopLevelStatementBreakContinueRejected(t *testing.T) {
	if !parseCompilationHasError(`break;`) {
		t.Fatalf("expect parse errors for top-level break")
	}
	if !parseCompilationHasError(`continue;`) {
		t.Fatalf("expect parse errors for top-level continue")
	}
}
