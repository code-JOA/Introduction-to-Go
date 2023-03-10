package main

import "fmt"

func main() {
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	var vals2 [3]int = vals
	fmt.Println(vals, vals2)

	// # Same as the above

	var v [3]int
	v[0] = 2
	v[1] = 4
	v[2] = 6
	fmt.Println(v, v[0], v[1], v[2])

	// # next array

	arr := [4]int{1, 2, 3, 4}
	total := 0

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	fmt.Println(total)

	// Two dimentional arrays with 3 items inside
	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoD[0][2])

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

}
