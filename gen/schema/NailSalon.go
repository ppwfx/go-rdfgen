package schema

type NailSalonTrait struct {

}

type NailSalon struct {
	MetaTrait
	NailSalonTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewNailSalon() (x NailSalon) {
	x.Type = "http://schema.org/NailSalon"

	return
}
