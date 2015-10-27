package main

import "fmt"

func avg(nums...int)float64{
	sum := 0
	for n := range nums {
		sum += n		
	}	
	return float64(sum)/float64(len(nums))
}

func main(){
	f := avg(1,2,3,4)
	fmt.Println(f)
	
	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(avg(nums...))
}