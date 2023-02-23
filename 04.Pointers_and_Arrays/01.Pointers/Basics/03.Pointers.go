package main

import "fmt"

// func setTo10(a *int) {
// 	*a = 10
// }

// func main() {
// 	a := 20
// 	fmt.Println(a)
// 	setTo10(&a)
// 	fmt.Println(a)
// }

func new2() {
	a := new(int)

	fmt.Println(a)
	fmt.Println(*a)
}

func setTo10(a *int) {
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}