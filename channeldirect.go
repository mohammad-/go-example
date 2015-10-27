package main

import "fmt"

func ping(p chan <- string){
	fmt.Println("will ping....")
	p <- "Ping..."
	p <- "Ping..."
}

func pong(pi <- chan string, po chan <- string){
	fmt.Println("Waiting for ping....")
	po <- <- pi
	po <- <- pi
	
}

func main(){
	pi := make(chan string,2)
	po := make(chan string,2)
	
	go pong(pi, po)
	go ping(pi)
	
	fmt.Println(<- pi)		
	fmt.Println(<- pi)
}
