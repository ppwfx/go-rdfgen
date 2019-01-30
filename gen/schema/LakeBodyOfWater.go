package schema

type LakeBodyOfWaterTrait struct {

}

type LakeBodyOfWater struct {
	MetaTrait
	LakeBodyOfWaterTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewLakeBodyOfWater() (x LakeBodyOfWater) {
	x.Type = "http://schema.org/LakeBodyOfWater"

	return
}
