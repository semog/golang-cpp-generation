package main

import "fmt"

func generated() {
	lst := Cons_int(3)
	lst = lst.Cons(Cons_int(4))
	first := lst.Car()
	second := lst.Cdr().Car()
	fmt.Println(first)
	fmt.Println(second)
	fmt.Printf("Add: %d\n", first+second)
}


