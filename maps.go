package main

import "fmt"

func main(){
	a := make(map[string]int)
	fmt.Println(a)
	
	a["key1"] = 10
	fmt.Println(a)
	
	a["key2"] = 20
	fmt.Println(a)

	a["key3"] = 30
	fmt.Println(a)
	
	v := a["key1"]
	fmt.Println(v)

	p , x := a["key4"]
	if x {
		fmt.Println(p)	
	}else{
		fmt.Println(x)
	}
	
	q , y := a["key3"]
	if y {
		fmt.Println(q,y)	
	}

}