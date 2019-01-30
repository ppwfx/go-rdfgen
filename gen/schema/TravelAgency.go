package schema

type TravelAgencyTrait struct {

}

type TravelAgency struct {
	MetaTrait
	TravelAgencyTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTravelAgency() (x TravelAgency) {
	x.Type = "http://schema.org/TravelAgency"

	return
}
