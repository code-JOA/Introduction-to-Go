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
