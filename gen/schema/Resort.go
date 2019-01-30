package schema

type ResortTrait struct {

}

type Resort struct {
	MetaTrait
	ResortTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewResort() (x Resort) {
	x.Type = "http://schema.org/Resort"

	return
}
