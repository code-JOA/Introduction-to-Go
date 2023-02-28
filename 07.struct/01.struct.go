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
	fmt.Println(h.A)
	fmt.Printf("---------\n")


	f1 := Cart{
		A: 20,
	}
	var f2 Cart
	f2 = f1
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(f1.A)

	var f3 *Cart = &f1
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(f1.A)

}











