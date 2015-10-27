package main

import "fmt"
import "time"

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func(){
		time.Sleep(time.Second * 1)
		c1 <- "One"
	}()
	go func(){		
		c2 <- "Two"
		time.Sleep(time.Second * 1)
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2) 
		}		
	}			
}
