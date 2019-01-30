package schema

type PoliceStationTrait struct {

}

type PoliceStation struct {
	MetaTrait
	PoliceStationTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	EmergencyServiceTrait
	LocalBusinessTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPoliceStation() (x PoliceStation) {
	x.Type = "http://schema.org/PoliceStation"

	return
}
