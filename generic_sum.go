package main

import "fmt"

type Integer int

func (i Integer) Add(j Integer) Integer {
	return Integer(i + j)
}

func generic_sum() {
	lst := Cons(Int(1))
	lst = lst.Cons(Cons(Int(2)))
	fmt.Println(lst.Sum().(Int) + 1)
}

