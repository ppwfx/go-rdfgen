package schema

type MotelTrait struct {

}

type Motel struct {
	MetaTrait
	MotelTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMotel() (x Motel) {
	x.Type = "http://schema.org/Motel"

	return
}
