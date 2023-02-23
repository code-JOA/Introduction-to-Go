package main

import "fmt"

func main() {
	b := new(int)

	fmt.Println(b)
	fmt.Println(*b)
}

func setTo10(a *int) {
	*a = 10
}

func new() {
	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}
