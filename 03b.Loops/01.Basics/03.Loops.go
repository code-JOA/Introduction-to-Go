package main

import "fmt"

func main() {

	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + i
	}
	fmt.Println(i)
}

// # the above can also be written this way
// func new() {
// 	i := 1
// 	for {
// 		fmt.Println(i)
// 		i = i + i
// 		if i > 10 {
// 			break
// 		}
// 	}
// 	fmt.Println(i)

// }
