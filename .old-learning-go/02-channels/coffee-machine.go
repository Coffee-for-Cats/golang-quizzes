package coffee_machine

import (
	"fmt"
)

func main() {
	// it's possible to limite the size of the channel >->
	// if it's full, block until something "gets out"
	coffee_machine := make(chan coffee, 2)
	// if I wrote the function here, it would have no need to pass the channel as parameter
	go MakeCoffee(127, milk, coffee_machine) // yumi
	go MakeCoffee(56, black, coffee_machine) // wtf?

	fmt.Println("Making your coffee!")
	// "<-" 's operations are blocky, but it's not lazy (yayy!)
	fmt.Println("A nice coffee for you:", <-coffee_machine)
	fmt.Println("Another one, but it's bad:", <-coffee_machine)
}
