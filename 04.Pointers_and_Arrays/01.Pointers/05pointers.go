package main

import (
	"fmt"
)

func main() {
	var gamepoint float64 = 99.2
	gameshooterRef := &gamepoint

	lifebooster = gamepoint * 2.2

	fmt.Println(gamepoint)
	fmt.Println(*lifeboosterRef)

}
