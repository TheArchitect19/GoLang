package main

import (
	"fmt"
)

func main(){
	fmt.Println("Process Started")
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	fmt.Println("execution finshed")
	
}