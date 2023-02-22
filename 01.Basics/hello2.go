package main

import (
	"fmt"
)

func main() {
	s := "Hello, world!"
	s2 := "👋 🌍"
	s3 := s + " " + s2
	fmt.Println(s3)

	x := "Go to, \n\t\"the world!\" with a backslash \\"
	// fmt.Println(x)

	b := s[0]
	b2 := s[4]
	fmt.Println(x, b, string(b), b2, string(b2), string(b2), string(b))

}

// func main() {
// 	s := "Hello, world!"
// 	b := s[0]
// 	b2 := s[4]
// 	fmt.Println(s, b, string(b), b2, string(b2))
// }
