package main

import (
	"fmt"
)

func main() {
	var isLoggedIn bool = true
	// var isLoggedIn bool = false
	var balance int = 10

	// note else should start right after the if stmt ends and not on the next line or error
	// && means and "and" || means Or
	if isLoggedIn && balance > 15 {
		fmt.Println("Show cart page")
	} else {
		fmt.Println("Show user login page")
	}
}
