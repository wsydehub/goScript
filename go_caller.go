package goScript

import "reflect"

// GoCaller is a thin adapter for calling methods on Go structs.
type GoCaller struct {
	structVal  reflect.Value
	structType reflect.Type
}

// CallFunc invokes a method on the underlying Go struct by name.
func (t *GoCaller) CallFunc(name string, params []*Variable) {

}
