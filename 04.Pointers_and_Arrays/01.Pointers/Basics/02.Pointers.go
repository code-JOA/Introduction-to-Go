// **********************************************Warning *****************************************************************

// Note: The code below will gennerate an error since we didnt assign any variable to b

// **********************************************Warning *****************************************************************

package main

import "fmt"

func main() {
	var b *int

	fmt.Println(b)
	fmt.Println(*b)
}


// This will also cause an error because nothing has been assigned to 'b'

func main() {
	var b : int()
	fmt.Println(b)
}