package goScript

import "reflect"

type GoCaller struct {
	structVal  reflect.Value
	structType reflect.Type
}

func (t *GoCaller) CallFunc(name string, params []*Variable) {

}
