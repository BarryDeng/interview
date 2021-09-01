package main

import (
	"fmt"
)

func main() {
	fmt.Println([...]string{"1"} == [...]string{"1"})
	// 切片不可直接比较
	//fmt.Println([]string{"1"} == []string{"1"})
}
