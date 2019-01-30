package schema

type FireStationTrait struct {

}

type FireStation struct {
	MetaTrait
	FireStationTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	EmergencyServiceTrait
	LocalBusinessTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFireStation() (x FireStation) {
	x.Type = "http://schema.org/FireStation"

	return
}
