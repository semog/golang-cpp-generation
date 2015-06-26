# Golang generics via code generation using the C preprocessor

21 Jun 2015 (by adityadodbole)

When Go came out, it was positioned as a better C++. However, Go does not have generics. And interfaces are not generics, no matter what anyone says. So I don’t see people who use C++ templates, moving to Go. And by that I don’t mean that I don’t think they won’t move; I have asked a lot of people who write C++ and they have said they are not moving because they need generics.

After being heavily influenced by [SICP](http://mitpress.mit.edu/sicp/) and having used dynamically typed languages like Ruby and Javascript, and to a lesser extent Java 8 and Haskell for the last 5 odd years, I was quite disappointed that there was no simple way to express higher order process abstractions in Go. Sure it has first class functions. But in a statically typed language, they are not very useful without generics or algebraic data types. There are a couple of options to solve the generics problem though. The first one is to use reflection. Here are a couple of files that demonstrate how to implement a “generic” cons list using reflection.

**reflected.go**
```
package main

import "fmt"

// Demo code using the definitions in reflect_cons.go  
// Notice the type conversions  
// Also the list is not type safe. One can even do  
// `lst = lst.Cons(Cons("imastring"))` and the compiler  
// wont complain and you will get a runtime panic  
func reflected() {
  lst := Cons(1)
  lst = lst.Cons(Cons(2))
  first := lst.Car().(int)
  second := lst.Cdr().Car().(int)
  fmt.Println(first)
  fmt.Println(second)
  fmt.Printf("Add: %d\n", first+second)
}

type Integer int

// Have to create a wrapper type with a Add method
// since reflect package cannot be used to call
// operators like "+"
func (i Integer) Add(j Integer) Integer {
	return Integer(i + j)
}

func reflected_sum() {
	lst := Cons(Integer(1))
	lst = lst.Cons(Cons(Integer(2)))
	fmt.Println(lst.Sum().(Integer) + 1)
}
```

**reflect_cons.go**
```
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
```
The above implementation has the following problems

* It is not typesafe
* Using `interface{}` is akin to using `void *` in C. You need to explicitly typecast the values
* There is a runtime hit. In an experiment involving creation of an actor using reflection, I got a 13% overhead compared to an implementation using a dispatch table.
* Implementing something as simple as a recursive sum is a PITA.

The other option is to implement something equivalent to generics using code generation. The options I came across were using the ast package or using go gen. The AST option was just too low level for me. The go gen packages seems good, but there was no documentation regarding writing your own typewriters and overall I didn’t think it was simple.

### The cpp solution

Initially I was attracted to Go because it is very similar to C in that it is simple, and has facilities that cater to fundamental orthogonal concepts of programming languages. So I had a hope that at least, Go could be a better C, at least where having a Go runtime is possible.

Recently, I happened to refer to some code I had written a long time ago in C and low and behold, I realised that I had solved the generics problem 10 years ago in C using code generation. In C, it’s called macros! So I tried using the C pre-processor on go code interspersed with some C preprocessor (cpp) code. And it worked, more or less. One issue was that cpp removes all newlines, which is unacceptable in go. To get around this, I terminated all lines with `;\` instead of the usual `\` and wrote a script that calls the C preprocessor and pipes the result through a sed command that replaces `;` with newlines.

The biggest advantage of this approach is that a lot of programmers have experience with the C preprocessor and understand it. Also, the whole approach is very simple, almost naive.

Here is the result for the above cons list implementation using type specialised code generation.

#### The macro code

**pair.h**
```
// The macro definitions for generating type specific code
// Since the C preprocessor removes newline, we need to end every line with
// ";\". The gopp script will replace the ";" with a newline.

// The CONSTYPE macro enforces the type naming convention
#define CONSTYPE(T) ConsP_##T

#define MAKE_CONS(T) ;\
	type CONSTYPE(T) struct {;\
		car T;\
		cdr *CONSTYPE(T);\
	};\
   	func Cons_##T (t T) *CONSTYPE(T){;\
		return &CONSTYPE(T) {;\
			car: t,;\
		};\
	};\
	func (p *CONSTYPE(T)) Car() T {;\
		return p.car;\
	};\
	func (p *CONSTYPE(T)) Cdr() *CONSTYPE(T) {;\
		return p.cdr;\
	};\
	func (p *CONSTYPE(T)) Cons(n *CONSTYPE(T)) *CONSTYPE(T) {;\
		n.cdr = p;\
		return n;\
	};
```

The Go code that will pass through the pre-processor

**pair.go.H**
```
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
```

The Go driver code

**generated.go**
```
package main

import "fmt"

// This calls the "generic" cons implementation for int
// generated by running cpp on pair.go.H
// No type conversions and type safe
func generated() {
	lst := Cons_int(3)
	lst = lst.Cons(Cons_int(4))
	first := lst.Car()
	second := lst.Cdr().Car()
	fmt.Println(first)
	fmt.Println(second)
	fmt.Printf("Add: %d\n", first+second)
}

type Int int

// One can just use the + operator here, but I'm making it
// as comparable to the version using reflection
func (i Int) Add(j Int) Int {
	return i + j
}

func generated_sum() {
	lst := Cons_Int(Int(1))
	lst = lst.Cons(Cons_Int(Int(2)))
	fmt.Println(lst.Sum() + 1)
}
```

The gopp (go pre-processor) script
```
#!/usr/bin/env bash

dir=$1
if [[ -z $dir ]]; then
	dir=.
fi
pushd $dir &>/dev/null
for i in `ls *.go.H`; do
	name=`basename $i .H`
	outfile="`basename $name .go`_gen.go"
	ifile=".$name.goi"
# Run the file through the C pre-proecssor and replace trailing ";" with
# newlines
	cc -E -P $i | sed -e $'s/;/\\\n/g' > $ifile
	gofmt $ifile > $outfile
	rm $ifile
done
popd &>/dev/null
```

### A few notes

* Only files ending with `.go.H` get acted upon by the gopp script. The code generated for a file named `foo.go.H` will be `foo_gen.go`.
* The generic type that is generated is named as `Foo_T` where `T` is the type for which the code is specialised. Eg. if in Java one would write `Cons<int>`, here it would be `Cons_int`.
* This approach doesn’t really solve the problem of creating libraries that implement generic types or algorithms (unless they are distributed as `.h` files), but it is a decent start
