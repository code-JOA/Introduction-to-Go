package main

import "fmt"

func main() {
	if age >= 18 {
		fmt.Println("You can drive alone")
	} else if age >=14 && age<18 {
		fmt.Println("You can drive with a Parent")
	} else if age < 14 {
		fmt.Println("You cannot drive. You are inderage")
	}
}
