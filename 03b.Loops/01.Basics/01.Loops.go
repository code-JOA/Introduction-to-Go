package main

import "fmt"

func main() {
	start := 1
	things := []string{"suitcase", "vase", "flower", "house", "car"}

	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	// # shorter syntax
	for i := range things {
		fmt.Println(things[i])
	}

	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
	// }

	// # There are no while loops in Go but we can use a similar syntax

	for start < 100 {
		start += start
		if start == 32 {
			break
		}
		fmt.Println("Start is now: ", start)
	}

	for start < 100 {
		start += start
		if start == 32 {
			continue
		}
		fmt.Println("Start is now: ", start)
	}

	for start < 100 {
		start += start
		if start == 32 {
			goto mylabel
		}
		fmt.Println("Start is now: ", start)

	mylabel:
		fmt.Println("I can Go")
	}

}
