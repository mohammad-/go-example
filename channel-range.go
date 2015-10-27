package main

import "fmt"

func main(){
	c := make(chan string, 2)
	
	c <- "One"
	c <- "Two"
	
	close(c)
	
	for msg := range c {
		fmt.Println(msg)
	}
}