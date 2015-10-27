package main

import "fmt"
// import "time"
func main(){
	c1 := make(chan string)

	go func(){
		fmt.Println("Waiting....")
		fmt.Println(<-c1)
	}()
	
	select {
		case msg := <- c1:
			fmt.Println(msg)
		default:
			fmt.Println("No message")
	}
	// time.Sleep(time.Second * 2)	
	select {
		case c1 <- "Yo":
			fmt.Println("message sent")
		default:
			fmt.Println("no message sent")
	}
}