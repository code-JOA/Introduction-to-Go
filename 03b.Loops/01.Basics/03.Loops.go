package main

import "fmt"



// # the above can also be written this way
func new2() {
	i := 1
	for {
		fmt.Println(i)
		i = i + i
		if i > 20 {
			break
		}
	}
	fmt.Println(i)
}
