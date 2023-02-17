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
	fmt.Println("Thor rating is: ", thorRating)

	// note printf is for formating options eg placeholders
	// "%v,%T" this %v shows the value and the %T shows the datatype( int or float)
	thorSpeed := 3.0
	fmt.Printf("%v,%T", "Thor Speed is:", thorSpeed, thorSpeed)

	var ironman, captAmerica string = "I am Ironman", "I am Capt America"
	fmt.Println(ironman, captAmerica)

	var defaultValue int
	fmt.Println(defaultValue)

	var(
		spiderman = "I am Spiderman", "I am Capt America"
	)
}
