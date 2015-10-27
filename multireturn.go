package main

import "fmt"

func add_and_sub(a int, b int)(int, int){
	x := a+b
	y := a-b
	return x,y
}

func s()(int, int){
	return 2,3
}
func main(){
	v1, v2 := s()
	fmt.Println(v1)
	fmt.Println(v2)	
	v3,v4 := add_and_sub(15, 10)
	fmt.Println(v3)
	fmt.Println(v4)
}