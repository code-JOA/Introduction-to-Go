package main

import "fmt"

func main() {
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

func m() {
	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
