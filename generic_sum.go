package main

import "fmt"

type Integer int

func (i Integer) Add(j Integer) Integer {
	return Integer(i + j)
}

func generic_sum() {
	lst := Cons(Integer(1))
	lst = lst.Cons(Cons(Integer(2)))
	fmt.Println(lst.Sum().(Integer) + 1)
}

