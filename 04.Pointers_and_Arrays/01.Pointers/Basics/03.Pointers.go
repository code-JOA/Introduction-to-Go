package main

import "fmt"

// Assignment 1
func new2() {
	a := new(int)

	fmt.Println(a)
	fmt.Println(*a)
}

// # Assignment 2
func setTo10(b *int) {
	*b = 10
}

func main() {
	b := 20
	fmt.Println(b)
	setTo10(&b)
	fmt.Println(b)
}


