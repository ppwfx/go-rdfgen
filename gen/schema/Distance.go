package schema

type DistanceTrait struct {

}

type Distance struct {
	MetaTrait
	DistanceTrait
	QuantityTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDistance() (x Distance) {
	x.Type = "http://schema.org/Distance"

	return
}
