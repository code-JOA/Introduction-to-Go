package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

// # print in increment of 2 Hence 2 , 4, 6, 8....

func new() {
	for x := 0; x < 10; x += 2 {
		fmt.Println("***********")
		fmt.Println(x)
	}
}
