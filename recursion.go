package main

import "fmt"

func findFact(n int)int{
	if(n == 0){
		return 1
	}else{
		return n * findFact(n - 1)
	}
}

func main(){
	fmt.Println(findFact(5))
}