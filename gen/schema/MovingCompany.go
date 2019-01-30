package schema

type MovingCompanyTrait struct {

}

type MovingCompany struct {
	MetaTrait
	MovingCompanyTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMovingCompany() (x MovingCompany) {
	x.Type = "http://schema.org/MovingCompany"

	return
}
