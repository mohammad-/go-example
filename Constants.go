package main

import "fmt"
import "math"

const s = "string"
func main(){
	fmt.Println(s)

	const pi = 22.0/7
	const largeN = 3e10
	fmt.Println(pi)
	
	fmt.Println(int64(largeN))
	
	const d = largeN / 3000
	fmt.Println(d)
	
	fmt.Println(math.Sin(d))

}