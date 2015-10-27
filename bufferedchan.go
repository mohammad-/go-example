package main

import "fmt"

func main(){
	message := make(chan rune,2)
	end := make(chan int)
	go func(){
		for true {
			c := <- message
			fmt.Printf("%c",c)
			if c == 0{
				fmt.Println()
				end <- 0
				break
			}	
		}		
	}()
	
	go func(){
		for _, c := range "Once upon a time, lived a king so great and nobel" {
			message <- c
		}		
		message <- 0
	}()
	
	<- end
}