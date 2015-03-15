package main

// This code is an implementation of a Cons pair and a "generic" Sum
// method using reflection
import "reflect"

type ConsP struct {
	car interface{}
	cdr *ConsP
}

func Cons(t interface{}) *ConsP {
	return &ConsP{
		car: t,
	}
}
func (p *ConsP) Car() interface{} {
	return p.car
}
func (p *ConsP) Cdr() *ConsP {
	return p.cdr
}
func (p *ConsP) Cons(n *ConsP) *ConsP {
	n.cdr = p
	return n
}

// Sum of list using recursive sum algorithm
// lst.sum = lst.car + lst.cdr.sum
func (p *ConsP) Sum() interface{} {
	v := reflect.ValueOf(p.car) // Get Value from interface{}
	t := reflect.TypeOf(p.car)  // Get Type from interface
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
	return val.Interface()      // Wrap in interface and return
}
