package main

import "fmt"

func main() {

	a := 3
	for i := 0; i < 10; i++ {
		if i == a {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println("Looping till it breaks", i)
	}
}
