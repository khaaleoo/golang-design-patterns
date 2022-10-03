package main

import (
	"fmt"

	abstract_factory "github.com/creational-patterns/abstract-factory"
	"github.com/creational-patterns/builder"
	factory_method "github.com/creational-patterns/factory-method"
	counter "github.com/creational-patterns/singleton"
)

func main() {
	/*
		Example Singleton
	*/
	fmt.Println("*** Example Singleton ***")
	// User does: Print, then increase counter...
	// At this time, we have no Counter instance yet.
	// So, we create one. With the value of 0.
	counter.GetInstance().Increase()

	// User prints again..., also increase counter...
	counter.GetInstance().Increase()

	fmt.Println("Print count: ", counter.GetInstance().Get())
	fmt.Print("*** End of Singleton ***\n\n\n")

	/*
		Example Builder
	*/
	fmt.Println("*** Example Builder ***")
	manufacturingVehicle := builder.ManufacturingDirector{}
	bicycleBuilder := &builder.BicycleBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.\n", bicycle.Structure, bicycle.Wheels, bicycle.Seats)
	fmt.Print("*** End of Builder ***\n\n\n")

	/*
		Example Factory Method
	*/
	fmt.Println("*** Example Factory Method ***")
	cappuccino, err := factory_method.GetCoffeeDrink("Cappuccino")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cappuccino.GetName())
	}

	latte, err := factory_method.GetCoffeeDrink("Latte")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(latte.GetName())
	}

	_, err = factory_method.GetCoffeeDrink("Error")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("*** End of Factory Method ***\n\n\n")

	/*
		Example Abstract Factory
	*/
	fmt.Println("*** Example Abstract Factory ***")

	bicycleFactory, err := abstract_factory.BuildFactory(abstract_factory.BicycleFactoryType)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	sportBicycle, err := bicycleFactory.NewVehicle(abstract_factory.SportBicycleType)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Sport Bicycle:")
	fmt.Printf("Vehicle has %d wheels, %d seats.\n", sportBicycle.NumWheels(), sportBicycle.NumSeats())

	fmt.Print("*** End of Abstract Factory ***\n\n\n")

}
