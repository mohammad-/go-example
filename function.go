package main

import "fmt"

func add(a int, b int)int{
	return a + b
}

func add3(a, b, c int)int{
	return a + b + c
}

func somef(a, b int, c string){
	fmt.Printf(c, a, b, a + b)
}

func main(){
	r := add(10, 5)
	fmt.Println(r)
	
	v := add3(10, 5, 10)
	fmt.Println(v)

	somef(10, 12, "%d + %d = %d")
}