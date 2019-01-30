package schema

type RiverBodyOfWaterTrait struct {

}

type RiverBodyOfWater struct {
	MetaTrait
	RiverBodyOfWaterTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewRiverBodyOfWater() (x RiverBodyOfWater) {
	x.Type = "http://schema.org/RiverBodyOfWater"

	return
}
