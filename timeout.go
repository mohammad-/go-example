package main

import "fmt"
import "time"

func func1(f1 chan <- string){
	time.Sleep(time.Second * 3)
	f1 <- "message 1"
}

func func2(f1 chan <- string){
	time.Sleep(time.Second * 1)
	f1 <- "message 2"
}


func main(){
	f1 := make(chan string)
	f2 := make(chan string)
	
	go func1(f1)
	go func2(f2)
	
	select {
		case m1 := <- f1:
			fmt.Println(m1)
		case <-time.After(time.Second * 1):
			  fmt.Println("time out for channel")
	}
	
	select {
		case m1 := <- f2:
			fmt.Println(m1)
		case <-time.After(time.Second * 1):
			  fmt.Println("time out for channel")
	}

}