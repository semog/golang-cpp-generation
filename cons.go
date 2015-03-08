package main

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

