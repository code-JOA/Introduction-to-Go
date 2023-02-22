package main

import (
	"fmt"
)

func main() {
	var numbers [3]string
	numbers[0] = "bread"
	numbers[1] = "chocolate"
	numbers[2] = "pizza"

	fmt.Println(numbers)

	var colors = [4]string{"rojo", "gris", "azul", "verde"}
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))

}
