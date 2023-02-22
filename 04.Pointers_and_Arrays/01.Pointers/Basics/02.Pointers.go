package main

import "fmt"

func main() {
	var b *int

	fmt.Println(b)
	fmt.Println(*b)
}


Note the above will gennerate an error since we didnt allow