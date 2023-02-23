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

	// print out position range 1 and take out the last item in the Array
	luggage = append(luggage[1 : len(luggage)-1])
	fmt.Println(luggage)
}
