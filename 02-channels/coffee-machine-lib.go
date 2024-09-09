package coffee_machine

import (
	"fmt"
	"time"
)

// enums are pretty weak typed? Strangy
type coffee_type int

const (
	black = iota
	milk
)

type coffee struct {
	temp        float32
	coffee_type coffee_type
}

// this way it implements the Stringer interface, so it prints beautifuly!
func (c coffee) String() string {
	milk_str := ""
	if c.coffee_type == milk {
		milk_str += "milk"
	} else {
		milk_str += "black"
	}

	return fmt.Sprintf("☕ - %s", milk_str)
}

// channels are op as heck
func MakeCoffee(temp float32, c_type coffee_type, send chan<- coffee) {
	fmt.Println("Coffee Machine working!")
	time.Sleep(time.Second * 2)
	send <- coffee{
		// must be explicit T-T
		temp:        temp,
		coffee_type: c_type,
	}
}
