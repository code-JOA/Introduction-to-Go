package main

import "fmt"

func main() {
	var a = 43
	b := &a
	b := c

	a = 10
	fmt.Println(a, b, *b)

	b = 20
	fmt.Println(b, *b, c)

	c = 40
	fmt.Println(b, *c)
}
