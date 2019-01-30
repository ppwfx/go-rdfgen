package schema

type HVACBusinessTrait struct {

}

type HVACBusiness struct {
	MetaTrait
	HVACBusinessTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHVACBusiness() (x HVACBusiness) {
	x.Type = "http://schema.org/HVACBusiness"

	return
}
