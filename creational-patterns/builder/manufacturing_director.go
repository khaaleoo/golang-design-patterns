package builder

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
