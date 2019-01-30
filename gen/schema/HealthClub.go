package schema

type HealthClubTrait struct {

}

type HealthClub struct {
	MetaTrait
	HealthClubTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	SportsActivityLocationTrait
	AdditionalTrait
}

func NewHealthClub() (x HealthClub) {
	x.Type = "http://schema.org/HealthClub"

	return
}
