package schema

type AdultEntertainmentTrait struct {

}

type AdultEntertainment struct {
	MetaTrait
	AdultEntertainmentTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAdultEntertainment() (x AdultEntertainment) {
	x.Type = "http://schema.org/AdultEntertainment"

	return
}
