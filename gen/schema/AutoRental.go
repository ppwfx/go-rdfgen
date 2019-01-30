package schema

type AutoRentalTrait struct {

}

type AutoRental struct {
	MetaTrait
	AutoRentalTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutoRental() (x AutoRental) {
	x.Type = "http://schema.org/AutoRental"

	return
}
