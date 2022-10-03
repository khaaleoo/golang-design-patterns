package abstract_factory

type NormalBicycle struct{}

func (*NormalBicycle) GetType() string {
	return "Normal Bicycle"
}

func (*NormalBicycle) NumWheels() int {
	return 2
}

func (*NormalBicycle) NumSeats() int {
	return 1
}
