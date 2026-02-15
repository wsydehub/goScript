package goScript

import "reflect"

// GoCaller is a thin adapter for calling methods on Go structs.
type GoCaller struct {
	structVal reflect.Value
}

func NewGoCaller(val interface{}) *GoCaller {
	if val == nil {
		return nil
	}
	rv := reflect.ValueOf(val)
	return &GoCaller{structVal: rv}
}

func (t *GoCaller) CallFunc(name string, params []interface{}) (interface{}, bool) {
	if t == nil || !t.structVal.IsValid() {
		return nil, false
	}
	rv := t.structVal
	method := rv.MethodByName(name)
	if !method.IsValid() && rv.Kind() != reflect.Ptr && rv.CanAddr() {
		addr := rv.Addr()
		method = addr.MethodByName(name)
	}
	if !method.IsValid() {
		return nil, false
	}
	var args []reflect.Value
	if params != nil {
		args = make([]reflect.Value, 0, len(params))
		for i, p := range params {
			rv := reflect.ValueOf(p)
			if i < method.Type().NumIn() {
				param := method.Type().In(i)
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
	results := method.Call(args)
	if len(results) == 0 {
		return nil, true
	}
	if len(results) == 1 {
		return results[0].Interface(), true
	}
	out := make([]interface{}, 0, len(results))
	for _, r := range results {
		out = append(out, r.Interface())
	}
	return out, true
}
