package main

import (
	"fmt"
)

func main() {
	fmt.Println("Process Started")
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	fmt.Println("execution finshed")

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("loop finsihed")
	/**
		5
		4
		3
		2
		1
	**/
}
