package schema

type RVParkTrait struct {

}

type RVPark struct {
	MetaTrait
	RVParkTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewRVPark() (x RVPark) {
	x.Type = "http://schema.org/RVPark"

	return
}
