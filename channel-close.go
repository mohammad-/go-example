package main

import "fmt"

func process(j <- chan string, done chan<- bool){
	for {
		select {
			case msg, more := <- j:
				if more {
					fmt.Println(msg)
				}else{
					fmt.Println("done")
					done <- true
					return					
				}
			default:
				fmt.Println("nothing")
					
		}
	}
}

func main(){
	done := make(chan bool)
	job := make(chan string, 5)
	go process(job, done)
	jobs := []string{"a","b","c","d"}
	for _, j := range(jobs){
		job <- j
	}	
	close(job)
	<- done 
}