package main

import "fmt"

func main(){
	var s = make([]string, 3)
	s[0] = "Mohammad"
	s[1] = "Sakina"
	s[2] = "Some one"
	fmt.Println(s)
	
	s = append(s,"x")
	s = append(s,"y")
	s = append(s,"z")
	s = append(s,"w")	
	fmt.Println(s)
	
	fmt.Println(s[3:])
	fmt.Println(s[3:4])
	fmt.Println(s[:3])
	
	t := []int{1,2,3,4}
	fmt.Println(t[1:3])			
	
	twoD := make([][]string, 5)
	for i:=0; i<5; i++{
		twoD[i] = make([]string,i)
		for j:=0; j<i; j++{		
			twoD[i][j] = "*"
		}
	}
	fmt.Println(twoD)
}