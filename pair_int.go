package main

// Implmentation of the Cons for int as a reference to implement
// the macro version. The macro version is derived by substituting
// the macro parameter T for int
/*
type ConsP_int struct {
	car int
	cdr *ConsP_int
}

func Cons_int(t int) *ConsP_int {
	return &ConsP_int{
		car: t,
	}
}
func (p *ConsP_int) Car() int {
	return p.car
}
func (p *ConsP_int) Cdr() *ConsP_int {
	return p.cdr
}
func (p *ConsP_int) Cons(n *ConsP_int) *ConsP_int {
	n.cdr = p
	return n
}
*/
