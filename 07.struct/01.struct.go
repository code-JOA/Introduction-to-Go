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


	i := Cart{
		A: 20,
	}
	var f2 Cart
	f2 = f
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(i.A)

	var f3 *Cart = &i
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(i.A)

}











