package schema

type SubwayStationTrait struct {

}

type SubwayStation struct {
	MetaTrait
	SubwayStationTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewSubwayStation() (x SubwayStation) {
	x.Type = "http://schema.org/SubwayStation"

	return
}
