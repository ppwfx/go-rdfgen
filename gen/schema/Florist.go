package schema

type FloristTrait struct {

}

type Florist struct {
	MetaTrait
	FloristTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFlorist() (x Florist) {
	x.Type = "http://schema.org/Florist"

	return
}
