package main

import "fmt"

func main() {
	var Batman string = "I am Batman"
	fmt.Println(Batman)

	var superman string = "I am Superman"
	var power int = 5
	fmt.Println(superman, "My power level is", power)

	var catwoman string
	catwoman = "I am catwoman"
	fmt.Println(catwoman)

	thor := "I am thor"
	fmt.Println(thor)
	thorRating := 3
	fmt.Println(thorRating)

	// note printf
	thorSpeed := 3
	fmt.Printf("%v, %T" , thorSpeed)

}
