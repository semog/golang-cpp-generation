package main

// Seperating the macros in a .h file so that they can become reusable
// Re-usable go libs as C header files! Oh the irony!
#include "pair.h"

// Generate the cons implementations for float64 and int
MAKE_CONS(float64)
MAKE_CONS(int)

// Define a macro to create a cons type with a Sum method
// Note that the sum method is defined only on this type. Other cons types
// are not affected.
// We use the CONSTYPE macro from pair.h to enforce type naming conventions
// for the generated code

#define MAKE_ADDER_CONS(T) MAKE_CONS(T);\
func (p *CONSTYPE(T)) Sum() T {;\
	if p.cdr == nil {;\
		return p.car;\
	};\
	return p.car.Add(p.cdr.Sum());\
}

MAKE_ADDER_CONS(Int)
