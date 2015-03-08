package main

import "fmt"

func generic() {
	lst := Cons(1)
	lst = lst.Cons(Cons(2))
	first := lst.Car().(int)
	second := lst.Cdr().Car().(int)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Printf("Add: %d\n", first+second)
}











// Not typesafe

