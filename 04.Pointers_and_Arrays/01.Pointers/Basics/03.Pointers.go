package main

import "fmt"

func new() {
	a := new(int)

	fmt.Println(a)
	fmt.Println(*a)
}

func setTo10(b *int) {
	*b = 10
}

func main() {
	b := 20
	fmt.Println(b)
	setTo10(&b)
	fmt.Println(b)
}

func setTo10Fail(c *int) {
	c = new(int)
	*c = 10
}

func new2() {
	c := 20
	fmt.Println(c)
	setTo10Fail(&c)
	fmt.Println(c)
}
