package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]
	fmt.Println("Existing array:\t", ref)
	t := append(s, ref...)
	fmt.Println("New slice:\t", t)
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)

	s = append(s, s...)
	fmt.Println("s+s:\t\t", s)

}

func new2() {
	s := make([]string, 0)
	fmt.Println("length of s:", len(s))
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	s2 := make([]string, 2)
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))
}
