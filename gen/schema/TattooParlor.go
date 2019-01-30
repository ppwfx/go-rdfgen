package schema

type TattooParlorTrait struct {

}

type TattooParlor struct {
	MetaTrait
	TattooParlorTrait
	HealthAndBeautyBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTattooParlor() (x TattooParlor) {
	x.Type = "http://schema.org/TattooParlor"

	return
}
