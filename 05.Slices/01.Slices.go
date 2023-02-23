package main

import (
	"fmt"
)

func main() {
	luggage := []string{"watch", "shirts", "trousers", "shoes"}
	fmt.Println(luggage)

	// let's add something else to my luggage
	luggage = append(luggage, "laptop")
	fmt.Println(luggage)
}
