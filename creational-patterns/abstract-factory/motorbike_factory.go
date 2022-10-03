package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	Motorbike125CCType = 1
	Motorbike150CCType = 2
)

type MotorbikeFactory struct{}

func (c *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case Motorbike125CCType:
		return new(Motorbike125CC), nil
	case Motorbike150CCType:
		return new(Motorbike150CC), nil
	default:
		err := fmt.Sprintf("Vehicle of type %d not recognized\n", v)
		return nil, errors.New(err)
	}
}
