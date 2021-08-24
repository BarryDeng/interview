package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan struct{}), make(chan struct{})
	wait := sync.WaitGroup{}
	wait.Add(1)

	go func() {
		i := 1
		for {
			<- number
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- struct{}{}
		}
	}()
	go func() {
		i := 'A'
		for {
			<- letter
			if i >= 'Z' {
				wait.Done()
				return
			}
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			number <- struct{}{}
		}
	}()

	number <- struct{}{}
	wait.Wait()
}
