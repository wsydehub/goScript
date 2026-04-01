package goScript

import "fmt"

// Function represents a script-level function definition.
type Function struct {
	Name           string
	ParametersList []*Variable
	ResultsList    []*Variable
	Scope          *Scope
	Block          *BlockContext
}

// NewFunction creates a function model with its scope and block body.
func NewFunction(name string, params, results []*Variable, scope *Scope, block *BlockContext) *Function {
	return &Function{
		Name:           name,
		ParametersList: params,
		ResultsList:    results,
		Scope:          scope,
		Block:          block,
	}
}

// InitParam binds call arguments into the function scope.
func (t *Function) InitParam(params []*Variable) error {
	if len(params) != len(t.ParametersList) {
		return fmt.Errorf("[InitParam] params's len is not equal")
	}

	for i, param := range t.ParametersList {
		inputParam := params[i]
		if param.TypeGo != inputParam.TypeGo {
			return fmt.Errorf("[InitParam] param type is not equal, expect: %v, actual: %v", param.TypeGo, inputParam.TypeGo)
		}
		param.Value = inputParam.Value
		t.Scope.AddVar(param)
	}

	return nil
}

// Exec runs the function body within its scope.
func (t *Function) Exec(visitor *Executor) {
	visitor.PushFunc(t)
	t.Block.Accept(visitor)
	visitor.PopFunc()
}

// SetResult assigns output values to declared result slots.
func (t *Function) SetResult(results []*Variable) error {
	if len(results) != len(t.ResultsList) {
		return fmt.Errorf("[SetResult] results's len is not equal")
	}

	for i, res := range t.ResultsList {
		inputRes := results[i]
		if res.TypeGo != inputRes.TypeGo {
			return fmt.Errorf("[SetResult] result type is not equal, expect: %v, actual: %v", res.TypeGo, inputRes.TypeGo)
		}
		res.Value = inputRes.Value
	}
	return nil
}
