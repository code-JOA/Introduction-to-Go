// package main

// import "fmt"

// func new2() {
// 	a := new(int)

// 	fmt.Println(a)
// 	fmt.Println(*a)
// }


// func setTo10(b *int) {
// 	*b = 10
// }

// func main() {
// 	b := 20
// 	fmt.Println(b)
// 	setTo10(&b)
// 	fmt.Println(b)
// }


package main

import "fmt"

func setTo10Fail(c *int) {
	a = new(int)
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)
}
