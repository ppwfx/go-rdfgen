package schema

type ComedyClubTrait struct {

}

type ComedyClub struct {
	MetaTrait
	ComedyClubTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewComedyClub() (x ComedyClub) {
	x.Type = "http://schema.org/ComedyClub"

	return
}
