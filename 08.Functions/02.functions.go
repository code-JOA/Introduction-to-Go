package main

import "fmt"

// add1 and add2 is just an explanation. we always return something when we write functions
func add1(x int, y int) {
	fmt.Println(x+y)
}

func add2(x, y int) {
	fmt.Println(x+y)
}

// Here we have two integers returning an int
func add3(x, y int) int {
	return x+y
}

// Here we have two separate integers returning two seperate integers, hence int int
//  also when we return two or more items we put them in brackets
func addandSub(x, y int) (int, int) {
	return x+y, x-y
}

func addandsub2(x, y int) (z1 int, z2 int) {
	z1 = x+y
	z2 = x-y
	return 
}

func deferstmt(x, y int) (x1, x2 int) {
	defer fmt.Println("Defer will return at the end of the function")
	x1 = x+y
	x2 = x-y
	fmt.Println("Proof i will print before defer")
	return
}

func main() {
	add1(20,1)
	add2(20,5)

	ans1, ans2 := addandSub(20, 7)
	fmt.Println(ans1, "\n", ans2)

	ans3, ans4 := addandsub2(20,10)
	fmt.Println(ans3, "\n", ans4)

	ans5, ans6 := deferstmt(30, 4)
	fmt.Println(ans5, "\n" , ans6)
}