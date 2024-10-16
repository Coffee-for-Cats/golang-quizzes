package main

import (
	"fmt"
	"time"
)

// interfaces are "generics" for other types and structs
type greenies interface {
	display()
	buy()
}

func buyTen(g greenies) {
	for i := 0; i < 10; i++ {
		g.buy()
	}
}

// fruit and it's implementations
type fruit struct {
	name string
	qtd  int
}

func (f *fruit) display() {
	fmt.Println("-", f.name, "\t |", f.qtd)
}
func (f *fruit) buy() {
	f.qtd++
}

// vegie and it's implementations
type vegie struct {
	name  string
	qtd   int
	price float32
}

// vegie also implements display
func (v *vegie) display() {
	fmt.Println(">", v.name, "\t |", v.qtd)
}

func (v *vegie) buy() {
	v.qtd++
}

func main() {
	// slice is basically a linked list in the end of the day
	buying_list := []greenies{}
	// reappend the pointer to new slice returned by append()
	// I need a reference here because the implementation of buy() only works in references.
	buying_list = append(buying_list, &fruit{name: "Banana"})
	buying_list = append(buying_list, &fruit{name: "Avocado"}) // 😈

	// // could panic if it's a vegie.
	// fmt.Println("I have", buying_list[0].(fruit).qtd, "bananas")

	// concurrency
	// prints the list every second, 3 times
	go func(fruits *[]greenies) {
		// executes 5 times
		for loops := 0; loops < 5; loops++ {
			// prints the whole list
			for i := 0; i < len(*fruits); i++ {
				(*fruits)[i].display()
				fmt.Println()
			}
			// waits 1 second.
			time.Sleep(time.Second)
		}
	}(&buying_list)

	// buys items in the main thread
	time.Sleep(time.Second)
	buying_list[0].buy()
	time.Sleep(time.Second * 2)
	buying_list[1].buy()
}
