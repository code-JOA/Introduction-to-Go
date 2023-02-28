package main 

import "fmt"

type Employee struct {
	Name string
	Email string
	age int
}

func main() {
	Josh := Employee {"Josh", "josh@gmail.com" , 23}
	fmt.Printf("%+v\n", Josh)
	fmt.Printf("%+v\n", Josh.Email)


	var Giulia = new(Employee)
	Giulia.Name = "Giulia"
	Giulia.Email = "giulia@gmail.com"
	Giulia.age = 22
	fmt.Printf("%+v\n", Giulia)

	var Siri = &Employee{"Siri", "siri@gmail.com", 29}
	fmt.Printf("%+v\n" , Siri)
}