package main

import (
	"fmt"
)




function main() {
	var isLoggedIn bool = true

	if isLoggedIn {
		fmt.Println("Show cart page")
	}
	else {
		fmt.Println("Show user login page")
	}
}