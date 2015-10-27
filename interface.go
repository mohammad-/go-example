package main

import "fmt"
import "math"

type geometry interface{
	area() float64
	prim() float64	
}

type rectangle struct{
	width float64
	height float64
} 

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (r rectangle) prim() float64{
	return (2 * r.width) + (2 * r.height)
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) prim() float64{
	return (2 * math.Pi * c.radius)
}

func printData(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.prim())
	
}

func main(){
	cir := circle{radius: 50}
	rect := rectangle{width:50, height:50}
	
	printData(cir)
	printData(rect)
}