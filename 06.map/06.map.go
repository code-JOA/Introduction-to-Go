package main

import "fmt"

func main() {
	balance := make(map[string]int)
	balance["Joshua"] = 89000000
	fmt.Println(balance)


	balance["James"] = 9000
	balance["Peter"] = 12000
	balance["John"] = 6000
	balance["Nebu"] = 5000

	getPeterBalance := balance["Peter"]
	fmt.Println(getPeterBalance)

	delete(balance, "Nebu")
	fmt.Println(balance)

	for key, value := range balance {
		fmt.Printf("balance of %v is %v\n", key, value)
	}

}
