package schema

type ParkingFacilityTrait struct {

}

type ParkingFacility struct {
	MetaTrait
	ParkingFacilityTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewParkingFacility() (x ParkingFacility) {
	x.Type = "http://schema.org/ParkingFacility"

	return
}
