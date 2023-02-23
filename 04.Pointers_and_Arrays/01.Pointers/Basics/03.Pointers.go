package main

import "fmt"

func main() {
	// a := 10
	// b := &a
	// c := b

	// a = 20
	// fmt.Println(a, b, c)

	// *b = 30
	// fmt.Println(a, *b, c)

	// c = 40
	// fmt.Println(a, b, *c)

	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

}
