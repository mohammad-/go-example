package main

import "fmt"

func main(){
	var a [5]int
	var b [10]string
	fmt.Println(a)
	fmt.Println(b)
	
	a[2] = 100
	b[3] = "Mohammad"
	
	fmt.Println(a)
	fmt.Println(b)

	a[1] = 20
	b[2] = "Sakina"

	fmt.Println(a)
	fmt.Println(b)
	x := [3]int{1,2,3}
	fmt.Println(x)
	
	var some2s [2][3]int
	some2s[1][1] = 10
	fmt.Println(some2s)
	
	for i:=0; i < 2; i++{
		for j:=0; j < 3; j++{
			some2s[i][j] = i + j
		}
	}
	fmt.Println(some2s)
}