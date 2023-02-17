package main

import (
	"fmt"
)

func main() {
	var isLoggedIn bool = true


note else should start right after the if stmt
	if isLoggedIn {
		fmt.Println("Show cart page")
	} else {
		fmt.Println("Show user login page")
	}
}
