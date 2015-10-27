package main

import ("fmt")

func main(){
	message := make(chan string)
	
	go func(){		
		message <- "How are you"
		fmt.Println("-------------")
		message <- "How are you"		
	}()
	
	fmt.Println("Waiting for message....")
	
	fmt.Println(<- message)	
	fmt.Println("*************")
	fmt.Println(<- message)
		
}