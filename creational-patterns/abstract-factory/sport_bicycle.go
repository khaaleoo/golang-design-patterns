package abstract_factory

type SportBicycle struct{}

func (*SportBicycle) GetType() string {
	return "Sport Bicycle"
}

func (*SportBicycle) NumWheels() int {
	return 1
}

func (*SportBicycle) NumSeats() int {
	return 1
}
