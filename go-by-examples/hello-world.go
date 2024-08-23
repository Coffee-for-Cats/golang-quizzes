package main

import "fmt"

func main() {

	for a := range 3 {
		if a%2 == 0 {
			fmt.Printf("even: %d\n", a)
		}
	}

	a := [...]int{1, 3: 7, 5, 4} // indexes 1 and 2 are 0
	for _, i := range a {
		if i == 7 {
			fmt.Println("Hello golang!")
		}
	}

	// this shit is powerfull
	compras := make([]string, 3) // no idea why it need to be "[]string"
	compras[0] = "batata"
	compras[1] = "banana"
	compras[2] = "beteraba"
	new_sliced := compras[1:3] // includes 1, don't includes 3, just like the default loop
	fmt.Println(new_sliced)
}
