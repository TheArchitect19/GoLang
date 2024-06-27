package main

import(
	"fmt"
)

func main(){
	fmt.Println("slices")
	s:= []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(s);

	fmt.Println(len(s))
	fmt.Println(cap(s))

	t:= s[1:4]
	fmt.Println(t)

	s= append(s, 69,35)
	fmt.Println(s)

	u:=make([]int, len(s))
	copy(u,s)
	fmt.Println(u)

	for i,v :=range u{
		fmt.Println(i," ",v)
	}
}