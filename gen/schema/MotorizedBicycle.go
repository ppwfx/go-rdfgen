package schema

type MotorizedBicycleTrait struct {

}

type MotorizedBicycle struct {
	MetaTrait
	MotorizedBicycleTrait
	VehicleTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewMotorizedBicycle() (x MotorizedBicycle) {
	x.Type = "http://schema.org/MotorizedBicycle"

	return
}
