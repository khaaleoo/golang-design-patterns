package abstract_factory

type Motorbike150CC struct{}

func (*Motorbike150CC) GetCC() string {
	return "150cc"
}

func (*Motorbike150CC) NumWheels() int {
	return 2
}

func (*Motorbike150CC) NumSeats() int {
	return 2
}
