package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		b := a / 2
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is less than or equal to 5:")
	}
}



func main() {
	a := 10
	if b := a / 2; b > 5 {
		fmt.Println("b is bigger than 5:", b)
	} else {
		fmt.Println("b is less than or equal to 5:", b)
	}
	fmt.Println(b)
}
