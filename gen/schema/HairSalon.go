package schema

type HairSalonTrait struct {

}

type HairSalon struct {
	MetaTrait
	HairSalonTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHairSalon() (x HairSalon) {
	x.Type = "http://schema.org/HairSalon"

	return
}
