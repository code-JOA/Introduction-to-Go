package main

import "fmt"

// func main() {
// 	b := new(int)

// 	fmt.Println(b)
// 	fmt.Println(*b)
// }

func setTo10(a *int) {
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}
