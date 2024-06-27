package main

import (
	"fmt"
)

func main() {
	fmt.Println("channels")

	var ch chan int
	ch = make(chan int)
	// ch1:= make(chan int)
	ch <- 42
	value := <-ch
	fmt.Println(value)

}
