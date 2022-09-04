package builder

type MotorbikeBuilder struct {
	v VehicleProduct
}

func (c *MotorbikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}
func (c *MotorbikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}
func (c *MotorbikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}
func (c *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}
