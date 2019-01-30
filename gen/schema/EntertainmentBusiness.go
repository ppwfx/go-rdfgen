package schema

type EntertainmentBusinessTrait struct {

}

type EntertainmentBusiness struct {
	MetaTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewEntertainmentBusiness() (x EntertainmentBusiness) {
	x.Type = "http://schema.org/EntertainmentBusiness"

	return
}
