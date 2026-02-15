package goScript

import "reflect"

// VariableType represents the internal runtime type category.
type VariableType int8

const (
	VarTypeInt     VariableType = 0
	VarTypeString  VariableType = 1
	VarTypeFloat   VariableType = 2
	VarTypeBool    VariableType = 3
	VarTypeChar    VariableType = 4
	VarTypeDynamic VariableType = 5
	VarTypeMap     VariableType = 6
	VarTypeArray   VariableType = 7
)

// Variable stores name, type metadata, and value.
type Variable struct {
	Name   string
	Type   VariableType
	TypeGo reflect.Type
	Value  reflect.Value
}

// NewVariable creates a variable with its runtime type and value.
func NewVariable(name string, t VariableType, typeGo reflect.Type, value reflect.Value) *Variable {
	return &Variable{
		Name:   name,
		Type:   t,
		TypeGo: typeGo,
		Value:  value,
	}
}
