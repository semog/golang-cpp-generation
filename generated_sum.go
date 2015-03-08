package main

import "fmt"

type Int int

func (i Int) Add(j Int) Int {
	return i + j
}

func generated_sum() {
	lst := Cons_Int(Int(1))
	lst = lst.Cons(Cons_Int(Int(2)))
	fmt.Println(lst.Sum() + 1)
}


