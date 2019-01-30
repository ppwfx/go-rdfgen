package schema

type ChildCareTrait struct {

}

type ChildCare struct {
	MetaTrait
	ChildCareTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewChildCare() (x ChildCare) {
	x.Type = "http://schema.org/ChildCare"

	return
}
