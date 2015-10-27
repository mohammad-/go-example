package main

import "fmt"

func f1(x int){
	x = 10
} 

func f2(x *int){
	*x = 10
} 


func main(){
	x := 5
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	f2(&x)
	fmt.Println(x)
	fmt.Println(&x)
}