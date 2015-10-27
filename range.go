package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5}
	for _,x := range a{
		fmt.Println(x)
	}
	
	for i,x := range a{
		fmt.Printf("index: %d, value: %d\n",i, x)
	}
	
	for _, c := range "Mohammad"{
		fmt.Printf("%c", c)
	}
	fmt.Println()
	
	m := map[string]string{"a":"apple", "b":"banana"}
	
	for k, v := range m{
		fmt.Printf("%s -> %s\n",k,v)
	}
}