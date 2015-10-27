package main

import "fmt"
import "time"

func main(){
	
	i := 1
	
	switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")			
	}
	
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }
	
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good Morning")
		case t.Hour() < 16:
			fmt.Println("Good Afternoon")	
		case t.Hour() <= 24:
			fmt.Println("Good Evening")	
		default: 									
	}
	
}