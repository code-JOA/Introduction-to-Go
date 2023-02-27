package main 

import ("fmt"
	)


type Cart struct {
    A int
	b string
}

func main() {
	f :=  Cart{}
	fmt.Println(f)

	g := Cart{10, "eggs"}
    fmt.Println(g)

	h := Cart{
		b: "oranges",
	}
	fmt.Println(h)

	h.A = 100
	fmt.Print(h.A)
}








