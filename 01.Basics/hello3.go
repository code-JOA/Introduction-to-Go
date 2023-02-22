package main

import (
	"fmt"
)

func main() {

	s := "Hello, world!"
	s2 := s[0:5]
	s3 := s[7:12]
	s4 := s[:5]
	s5 := s[7:]
	fmt.Println(s, s2, s3, s4, s5)

	fmt.Println(s, "Has ", len(s), "words.", s2, "Has ", len(s2), " words.", s3, "Has ", len(s3), "words.")

	x := "Golang to the, "
	var r rune = 127757
	x = x + string(r)
	fmt.Println(x)

}
