package main

import ("fmt" 
		"time")

func f(end chan bool){
	fmt.Println("............")
	time.Sleep(time.Second * 1)
	fmt.Println("............")
	end <- true
}

func main(){
	end := make(chan bool)
	go f(end)
	<-end
}