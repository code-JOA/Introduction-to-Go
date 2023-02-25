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

	// # Here we want to create heroes, with 3 elements with an upperbound of 3 but note we can add more items with slices
	mavel := make([]string, 3, 3)

	mavel[0] = "Batman"
	mavel[1] = "Superman"
	mavel[2] = "Spiderman"

	fmt.Println(mavel)

	mavel = append(mavel, "Robin")
	fmt.Println(mavel)
	// # finding how much more i can add to heroes
	fmt.Println(cap(mavel))
}
