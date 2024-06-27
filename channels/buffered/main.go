package main

import "fmt"


func main(){
	
	var ch chan int
	go func ()  {
		ch <- 42
	}()
	val := <-ch
	fmt.Println(val)

}