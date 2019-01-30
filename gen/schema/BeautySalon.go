package schema

type BeautySalonTrait struct {

}

type BeautySalon struct {
	MetaTrait
	BeautySalonTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBeautySalon() (x BeautySalon) {
	x.Type = "http://schema.org/BeautySalon"

	return
}
