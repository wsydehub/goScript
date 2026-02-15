package goScript

import (
	"reflect"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

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
