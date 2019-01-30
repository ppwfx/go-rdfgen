package schema

type NightClubTrait struct {

}

type NightClub struct {
	MetaTrait
	NightClubTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewNightClub() (x NightClub) {
	x.Type = "http://schema.org/NightClub"

	return
}
