package main

import (
	"fmt"
)

func main(){
	var a [3]int
	
	a=[3]int{1,2,3}
	b:=[3]int{4,5,6}
	fmt.Println(a[0])
	fmt.Println(b[0])

	for i:=0; i<len(a);i++ {
		fmt.Println(a[i])
	}
	for index, value := range b{
		fmt.Println(index);
		fmt.Print(" ",value)
	}

	matrix:= [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	for index :=range matrix {
		for i, val2 := range matrix[index] {
			fmt.Println(i)
			fmt.Println(val2)
		}
	}
}