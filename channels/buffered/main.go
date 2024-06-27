package main

import (
	"fmt"
)

func main(){
	ch:= make(chan int, 2)
	ch<-1
	ch<-2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)

	value, ok := <-ch
	if !ok{
		fmt.Println("Channel is closed")
	}else {
		fmt.Println(value)
	}
	for val := range ch {
		fmt.Println(val)
	}
}