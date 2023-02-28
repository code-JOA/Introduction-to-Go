package main 

import "fmt"


func main() {
	coder()
	result := multiplyme(3 , 9)
	fmt.Println(result)

	myresult := adder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(myresult)

	myresult2, mylength, myname := multiplyer2(1,2,3,4,5)
	fmt.Println(myresult2, mylength, myname)
}

func coder() {
	fmt.Println("I am learning to code in Golang")
}

func multiplyme(x1 int, x2 int) int {
	return x1 * x2
	// fmt.Println("This will execute after we use a return so lets harsh it out to prevent an error")
}

func adder(values ...int) int {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	return sum
}


func multiplyer2(values ...int) (int, int, string)	{
	multi := 1
	for x := range values {
		multi = multi * values[x]
	}
	length := len(values)
	name := "Go-coder"
	return multi, length, name



}
