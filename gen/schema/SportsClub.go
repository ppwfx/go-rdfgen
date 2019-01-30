package schema

type SportsClubTrait struct {

}

type SportsClub struct {
	MetaTrait
	SportsClubTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewSportsClub() (x SportsClub) {
	x.Type = "http://schema.org/SportsClub"

	return
}
