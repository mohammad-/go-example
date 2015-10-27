package main

import "fmt"

func intSeq(i int) func() int{
	
	return func() int{
		i = i + 1
		return i
	}
}

func main(){
	f := intSeq(5)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	
	f1 := intSeq(0)
	fmt.Println(f1())
}