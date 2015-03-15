package main

type ConsP_float64 struct {
	car float64
	cdr *ConsP_float64
}

func Cons_float64(t float64) *ConsP_float64 {
	return &ConsP_float64{
		car: t,
	}
}
func (p *ConsP_float64) Car() float64 {
	return p.car
}
func (p *ConsP_float64) Cdr() *ConsP_float64 {
	return p.cdr
}
func (p *ConsP_float64) Cons(n *ConsP_float64) *ConsP_float64 {
	n.cdr = p
	return n
}

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
