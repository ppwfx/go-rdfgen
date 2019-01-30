package schema

type HomeAndConstructionBusinessTrait struct {

}

type HomeAndConstructionBusiness struct {
	MetaTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHomeAndConstructionBusiness() (x HomeAndConstructionBusiness) {
	x.Type = "http://schema.org/HomeAndConstructionBusiness"

	return
}
