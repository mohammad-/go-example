package main

import "fmt" 
import "errors"


type err struct{
	message string
	code int
}

func (e *err)Error() string{
	return fmt.Sprintf("%d -- %s ",e.code, e.message)
}
		
type ride interface{
	drive()(string, error)
	fly()(string, error)
}

type person struct{
	age int
	name string
}

func (p person) drive()(string, error){
	if p.age < 18{
		return "", errors.New("You are small dude...")
	}else if p.age > 18 && p.age < 75{
		return "Take my car....", nil	
	}else{
		return "", errors.New("You need to get checked first...")
	}
}

func (p person) fly()(string, error){
	if p.age < 20{
		return "", &err{code: 0, message:"Too young to fly"}
	}else if p.age > 20 && p.age < 75{
		return "Take my jet....", nil	
	}else{
		return "", &err{code: 1, message:"Too old to fly"}
	}
}
		

func main(){
	fmt.Println("Lets check who can drive")
	p := []person{person{name:"Mohammad", age:29},
	 			  person{name:"someone", age:10}, 
				   person{name:"someone", age:90}}
				   		
	for _, per := range p{		
		if m ,e := per.drive(); e == nil{
			fmt.Printf("%s, %s\n",per.name, m)
		}else{
			fmt.Println(e)
		}
	}
	
	for _, per := range p{		
		if m ,e := per.fly(); e == nil{
			fmt.Printf("%s, %s\n",per.name, m)
		}else{
			fmt.Println(e)
		}
	}

	x := person{name:"some name", age:10}
	_, e := x.fly()
	if ex, ok := e.(*err); ok{
		fmt.Println(ex.code, ex.message)	
	}

	y := person{name:"some name", age:25}
	_, eex := y.fly()	
	if see,ok := eex.(*err); ok {
		fmt.Println(see.code,see.message)	
	}else{
		fmt.Println(ok)
	}
	
}		
