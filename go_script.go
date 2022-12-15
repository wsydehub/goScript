package goScript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GoScript struct {
	input    antlr.CharStream
	lexer    *GoScriptLexer
	tokens   *antlr.CommonTokenStream
	parser   *GoScriptParser
	tree     ICompilationUnitContext
	executor *Executor
}

func (t *GoScript) Init(input antlr.CharStream) error {
	t.lexer = NewGoScriptLexer(input)
	t.tokens = antlr.NewCommonTokenStream(t.lexer, antlr.TokenDefaultChannel)
	t.parser = NewGoScriptParser(t.tokens)
	t.tree = t.parser.CompilationUnit()
	t.executor = NewExecutor()
	t.tree.Accept(t.executor)
	return nil
}
