package main

import "fmt"

func main() {
	var b = 43
	c := &b
	*b := c

	fmt.Println(b)

	b = 20
	fmt.Println(b, *b)

	c = 40
	fmt.Println(*c)
}
