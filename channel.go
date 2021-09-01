package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		fmt.Printf("%p\n", ch)
		ch <- 1
		fmt.Println("OVER")
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		fmt.Printf("2: %p\n", ch)
		<-ch
		fmt.Printf("2: over\n")
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}