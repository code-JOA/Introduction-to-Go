package main

import "fmt"

func main() {
	age := 18
	can_drink = "True"

	if age >= 18 {
		fmt.Println("You can drive alone")
	} else if age >= 14 && age < 18 {
		fmt.Println("You can drive with a Parent")
	} else {
		fmt.Println("You cannot drive. You are underage")
	}
}
