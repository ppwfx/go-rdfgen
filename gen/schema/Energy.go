package schema

type EnergyTrait struct {

}

type Energy struct {
	MetaTrait
	EnergyTrait
	QuantityTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEnergy() (x Energy) {
	x.Type = "http://schema.org/Energy"

	return
}
