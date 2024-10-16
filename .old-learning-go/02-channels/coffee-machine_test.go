package coffee_machine

import (
	"testing"
)

func TestMakeCoffee(t *testing.T) {
	channel := make(chan coffee, 1)
	go MakeCoffee(100, black, channel)
	expected := coffee{
		temp:        100,
		coffee_type: black,
	}
	if <-channel != expected {
		t.Error("Coffee was made wrongly!")
	}
}
