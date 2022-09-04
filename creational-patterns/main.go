package main

import (
	"fmt"

	"github.com/creational-patterns/builder"
	counter "github.com/creational-patterns/singleton"
)

func main() {
	/*
		Example Singleton
	*/
	// User does: Print, then increase counter...
	// At this time, we have no Counter instance yet.
	// So, we create one. With the value of 0.
	counter.GetInstance().Increase()

	// User prints again..., also increase counter...
	counter.GetInstance().Increase()

	// fmt.Println("Print count: ", counter.GetInstance().Get())

	/*
		Example Builder
	*/
	manufacturingVehicle := builder.ManufacturingDirector{}
	bicycleBuilder := &builder.BicycleBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.", bicycle.Structure, bicycle.Wheels, bicycle.Seats)
}
