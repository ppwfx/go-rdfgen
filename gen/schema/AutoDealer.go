package schema

type AutoDealerTrait struct {

}

type AutoDealer struct {
	MetaTrait
	AutoDealerTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutoDealer() (x AutoDealer) {
	x.Type = "http://schema.org/AutoDealer"

	return
}
