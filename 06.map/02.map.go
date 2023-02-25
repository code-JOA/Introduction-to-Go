package main 

import "fmt"


func main() {
	grocerry := make(map[string]int) 
		grocerry["apple"] = 5
		
		cart := grocerry["apple"]
	    fmt.Println(grocerry)
		fmt.Println("apple in grocerry:" , cart)
		fmt.Println("a in grocerry:" , grocerry["a"])

		if v, ok := grocerry["apple"]; ok {
			fmt.Println("apple in grocerry", v)
		}

		grocerry["apple"] = 20
		// fmt.Println("apples in grocerry:", grocerry["apple"])
		fmt.Printf("%v apples in grocerry", grocerry["apple"])
}


