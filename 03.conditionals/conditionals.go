package main

import (
	"fmt"
)

func main() {
	var isLoggedIn bool = true
	// var isLoggedIn bool = false
	var balance int = 4

	// note else should start right after the if stmt ends and not on the next line or error
	// && means and "and" || means Or
	if isLoggedIn && balance > 15 {
		fmt.Println("Show cart page")
		var len, err = fmt.Println("Show cart page")
		if err == nil {
			fmt.Println("Passed test", len)
			// using placeholders v is verbs and T is the datatype
			fmt.Printf("Balance is %v, %T", len, len)
			fmt.Println(`This is a Backtick`)
		}
		// note else should start right after the if stmt ends and not on the next line or error
		// } else if isLoggedIn && balance < 5 {
		// 	fmt.Println("Check if you have money in your account")
	} else {
		fmt.Println("Show user login page")
	}
}
