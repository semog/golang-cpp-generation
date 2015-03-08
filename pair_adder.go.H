package main

#include "pair.h"

#define MAKE_ADDER_CONS(T) MAKE_CONS(T);\
func (p *CONSTYPE(T)) Sum() T {;\
	if p.cdr == nil {;\
		return p.car;\
	};\
	return p.car.Add(p.cdr.Sum());\
}

MAKE_ADDER_CONS(Int)
