package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	fmt.Println(person{name:"Mohammad", age:29})
	var p person = person{name:"Sakina", age:22}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)
	
	p.name = "Some Name"
	fmt.Println(p)
	
	sp := p
	sp.name = "x"
	fmt.Println(sp)
	fmt.Println(p)
	
	sp2 := &p
	sp2.name = "yy"
	fmt.Println(sp2.name)
	fmt.Println(p)

}
