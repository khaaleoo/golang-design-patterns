package abstract_factory

type Motorbike125CC struct{}

func (*Motorbike125CC) GetCC() string {
	return "125cc"
}

func (*Motorbike125CC) NumWheels() int {
	return 2
}

func (*Motorbike125CC) NumSeats() int {
	return 2
}
