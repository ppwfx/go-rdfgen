package schema

type MotorcycleDealerTrait struct {

}

type MotorcycleDealer struct {
	MetaTrait
	MotorcycleDealerTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMotorcycleDealer() (x MotorcycleDealer) {
	x.Type = "http://schema.org/MotorcycleDealer"

	return
}
