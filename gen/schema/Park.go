package schema

type ParkTrait struct {

}

type Park struct {
	MetaTrait
	ParkTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPark() (x Park) {
	x.Type = "http://schema.org/Park"

	return
}
