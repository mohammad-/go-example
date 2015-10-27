package main

import "fmt"
import "time"
	
func Printstr(str string){
	for _, s := range str {
		fmt.Printf("%c",s)
		time.Sleep(1)				
	}
	fmt.Println("")
}


func main(){
	Printstr("mohammad bharmal")
	
	go Printstr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	go Printstr("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	
	go func(someval string){
		for _, c := range someval{
			fmt.Printf("%c",c)
			time.Sleep(1)
		}
	}("ccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
	
	var line string	
	fmt.Scanln(&line)
	fmt.Println("Done")
}