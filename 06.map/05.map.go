package main

import (
	"fmt"
)

func main() {

	myMap := make(map[string]int)
	myMap["k1"] = 12
	myMap["k2"] = 13
	fmt.Println("myMap:", myMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}

	fmt.Println("anotherMap:", anotherMap)
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	_, ok := myMap["doesItExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
