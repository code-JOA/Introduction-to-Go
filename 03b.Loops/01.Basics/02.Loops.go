package main

import "fmt"

func main() {
	x := 0
	for x <= 5 {
		fmt.Println(x)
		x++
	}

	// # this is the same as the code above

	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	for x := 0; x <= 5
}