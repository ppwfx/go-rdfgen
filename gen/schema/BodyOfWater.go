package schema

type BodyOfWaterTrait struct {

}

type BodyOfWater struct {
	MetaTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBodyOfWater() (x BodyOfWater) {
	x.Type = "http://schema.org/BodyOfWater"

	return
}
