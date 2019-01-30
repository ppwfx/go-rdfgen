package schema

type MotorcycleTrait struct {

}

type Motorcycle struct {
	MetaTrait
	MotorcycleTrait
	VehicleTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewMotorcycle() (x Motorcycle) {
	x.Type = "http://schema.org/Motorcycle"

	return
}
