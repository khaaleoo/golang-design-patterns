package builder

type BicycleBuilder struct {
	v VehicleProduct
}

func (c *BicycleBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}
func (c *BicycleBuilder) SetSeats() BuildProcess {
	c.v.Seats = 1
	return c
}
func (c *BicycleBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Bicycle"
	return c
}
func (c *BicycleBuilder) GetVehicle() VehicleProduct {
	return c.v
}
