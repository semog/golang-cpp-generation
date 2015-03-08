package main

import "reflect"
// lst.sum = lst.car + lst.cdr.sum
func (p *ConsP) Sum() interface{} {
	v := reflect.ValueOf(p.car) // Get Value from interface{}
	t := reflect.TypeOf(p.car) // Get Type from interface
	// Exit condition
	if p.cdr == nil {
		return v.Interface()
	}
	// Fetch the method to call
	m, _ := t.MethodByName("Add")
	// Setup method args
	args := make([]reflect.Value, 2)
	args[0] = v
	args[1] = reflect.ValueOf(p.cdr.Sum())
	val := m.Func.Call(args)[0] // Call
	return val.Interface() // Wrap in interface and return
}


// 13% overhead
