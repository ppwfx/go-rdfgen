package schema

type SeaBodyOfWaterTrait struct {

}

type SeaBodyOfWater struct {
	MetaTrait
	SeaBodyOfWaterTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewSeaBodyOfWater() (x SeaBodyOfWater) {
	x.Type = "http://schema.org/SeaBodyOfWater"

	return
}
