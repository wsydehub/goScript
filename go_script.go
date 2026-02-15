package goScript

import (
	"github.com/antlr4-go/antlr/v4"
)

// GoScript wires lexer/parser/executor for a single run.
type GoScript struct {
	input    antlr.CharStream
	lexer    *GoScriptLexer
	tokens   *antlr.CommonTokenStream
	parser   *GoScriptParser
	tree     ICompilationUnitContext
	executor *Executor
}

// Init parses the input and executes the compilation unit.
func (t *GoScript) Init(input antlr.CharStream) error {
	t.lexer = NewGoScriptLexer(input)
	t.tokens = antlr.NewCommonTokenStream(t.lexer, antlr.TokenDefaultChannel)
	t.parser = NewGoScriptParser(t.tokens)
	t.tree = t.parser.CompilationUnit()
	t.executor = NewExecutor()
	t.tree.Accept(t.executor)
	return nil
}
