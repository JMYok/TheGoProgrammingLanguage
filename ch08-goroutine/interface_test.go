package ch08_goroutine

import (
	"fmt"
	"testing"
)

type Vehicle interface {
}

type Car struct {
}

func TestInterface(t *testing.T) {
	var car1 *Car
	fmt.Println("The first car is nil. ")
	car2 := car1
	fmt.Println("The second car is nil. ")
	var vehicle Vehicle = car2
	if vehicle == nil {
		fmt.Println("The vehicle is nil. ")
	} else {
		fmt.Println("The vehicle is not nil. ")
	}
}
