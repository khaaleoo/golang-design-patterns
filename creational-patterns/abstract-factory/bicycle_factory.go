package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	NormalBicycleType = 1
	SportBicycleType  = 2
)

type BicycleFactory struct{}

func (c *BicycleFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case NormalBicycleType:
		return new(NormalBicycle), nil
	case SportBicycleType:
		return new(SportBicycle), nil
	default:
		err := fmt.Sprintf("Vehicle of type %d not recognized\n", v)
		return nil, errors.New(err)
	}
}
