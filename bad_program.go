package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Param map[string]interface{}

type Show struct {
	Param
}

type Student struct {
	Name string
}

func decide(s interface{}) string {
	switch msg := s.(type) {
	case *Student:
		fmt.Println(msg.Name)
	case Student:
		fmt.Println(msg.Name)
	}
	return "FF"
}



func main() {
	s := new(Show)
	s.Param = make(map[string]interface{})
	s.Param["abc"] = "efd"
	fmt.Println(s)

	decide(Student{Name: "fsdafs"})

	var a int32 = 1
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			a ++
		}
		wg.Done()
	}()
	go func() {
		for {
			if atomic.CompareAndSwapInt32(&a, 1000, 99999) {
				fmt.Println("F")
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(a)
}
