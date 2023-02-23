package main

import "fmt"

func main() {
	var a = 10
	b := &a
	c := b

	a = 20
	fmt.Println(a, b, *b)

	b = 30
	fmt.Println(b, *b, c)

	c = 40
	fmt.Println(b, *c)
}
