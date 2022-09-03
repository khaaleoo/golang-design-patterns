package main

import (
	"fmt"

	counter "github.com/creational-patterns/singleton"
)

func main() {
	// User does: Print, then increase counter...
	// At this time, we have no Counter instance yet.
	// So, we create one. With the value of 0.
	counter.GetInstance().Increase()

	// User prints again..., also increase counter...
	counter.GetInstance().Increase()

	fmt.Println("Print count: ", counter.GetInstance().Get())

}
