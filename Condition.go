package main

import "fmt"

func main(){
	if 9%3 == 0{
		fmt.Println(9, "can be devided by",3)
	}else{
		fmt.Println(9, "cannot be devided by",3)
	}
	
	if 9%2 == 0{
		fmt.Println(9, "can be devided by",2)
	}else{
		fmt.Println(9, "cannot be devided by",2)
	}

	if result := 9%2; result == 0{
		fmt.Println(9, "can be devided by",2)
	}else{
		fmt.Println(9, "cannot be devided by",2)
	}

}