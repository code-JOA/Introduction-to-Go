package main

import (
	"fmt"
)

func main() {
	var menu [3]string
	menu[0] = "bread"
	menu[1] = "chocolate"
	menu[2] = "pizza"

	fmt.Println(menu)

	var restaurant = [4]string{"prawns", "chicken", "meat", "vegetables"}
	fmt.Println(restaurant)
	fmt.Println(restaurant[2])
	fmt.Println(len(restaurant))

	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	var vals2 [4]int = vals
	fmt.Println(vals, vals2)

}
