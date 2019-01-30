package schema

type OceanBodyOfWaterTrait struct {

}

type OceanBodyOfWater struct {
	MetaTrait
	OceanBodyOfWaterTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewOceanBodyOfWater() (x OceanBodyOfWater) {
	x.Type = "http://schema.org/OceanBodyOfWater"

	return
}
