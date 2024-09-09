package main

import (
	"04-modules/package2"
	"04-modules/package2/package3"
	"fmt"
)

func main() {
	package2.SayHello()
	numero := package3.RandomInt()
	fmt.Println("Joguei um dado aqui: ", numero)
}
