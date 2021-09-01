package main

import (
	"container/list"
	"fmt"
)

type mylist struct {
	V int
	Next *mylist

}

func reverse(r *list.List)

func main() {
	r := &mylist{}
	p := r
	for i := 0; i < 10; i++ {
		p.V = i
		p.Next = &mylist{}
		p = p.Next
	}
	fmt.Println(r)
	reverse(r)
	fmt.Println(r)

}
