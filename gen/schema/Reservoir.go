package schema

type ReservoirTrait struct {

}

type Reservoir struct {
	MetaTrait
	ReservoirTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewReservoir() (x Reservoir) {
	x.Type = "http://schema.org/Reservoir"

	return
}
