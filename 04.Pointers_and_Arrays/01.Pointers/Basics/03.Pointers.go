package main

import "fmt"

func main() {
	var b int() = 43
	c := &b
	*b := c

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*c)
}
