package main

import (
	"fmt"
)

func main() {
	var gamepoint float64 = 99.2
	gameshooterRef := &gamepoint

	gamepoint = gamepoint * 5.2

	fmt.Println(gamepoint)
	fmt.Println(*gameshooterRef)

}
