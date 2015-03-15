package main

import "fmt"

// Demo code using the definitions in reflect_cons.go
// Notice the type conversions
// Also the list is not type safe. One can even do
// `lst = lst.Cons("imastring")` and the compiler
// wont complain and you will get a runtime panic
func generic() {
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

func generic_sum() {
	lst := Cons(Integer(1))
	lst = lst.Cons(Cons(Integer(2)))
	fmt.Println(lst.Sum().(Integer) + 1)
}

