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

	for x := 0; x <= 10; x += 2 {
		fmt.Println(x)
	}

	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
		}
	}

}
