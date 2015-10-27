package main

import "fmt"

type person struct{
	name string
	age int
}

func (p person) printname(){
	fmt.Println(p.name)
}

func (p person) issenior()bool{	
	return p.age > 60
}

func (p *person) addtoage(val int){
	p.age += val
}

func main(){
	p1 := person{name:"mohammad", age:29}
	p2 := person{name:"someone", age:90}
	p1.printname()
	p2.printname()
	
	
	fmt.Println(p1.issenior())
	fmt.Println(p2.issenior())
	
	fmt.Println(p2.age)	
	p2.addtoage(10)
	fmt.Println(p2.age)
} 