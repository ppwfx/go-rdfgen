package schema

type PondTrait struct {

}

type Pond struct {
	MetaTrait
	PondTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPond() (x Pond) {
	x.Type = "http://schema.org/Pond"

	return
}
