package schema

type HealthAndBeautyBusinessTrait struct {

}

type HealthAndBeautyBusiness struct {
	MetaTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHealthAndBeautyBusiness() (x HealthAndBeautyBusiness) {
	x.Type = "http://schema.org/HealthAndBeautyBusiness"

	return
}
