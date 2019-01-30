package schema

type MassTrait struct {

}

type Mass struct {
	MetaTrait
	MassTrait
	QuantityTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMass() (x Mass) {
	x.Type = "http://schema.org/Mass"

	return
}
