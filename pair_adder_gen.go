package main

type ConsP_Int struct {
	car Int
	cdr *ConsP_Int
}

func Cons_Int(t Int) *ConsP_Int {
	return &ConsP_Int{
		car: t,
	}
}
func (p *ConsP_Int) Car() Int {
	return p.car
}
func (p *ConsP_Int) Cdr() *ConsP_Int {
	return p.cdr
}
func (p *ConsP_Int) Cons(n *ConsP_Int) *ConsP_Int {
	n.cdr = p
	return n
}

func (p *ConsP_Int) Sum() Int {
	if p.cdr == nil {
		return p.car
	}
	return p.car.Add(p.cdr.Sum())
}
