package schema

type AttorneyTrait struct {

}

type Attorney struct {
	MetaTrait
	AttorneyTrait
	LegalServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAttorney() (x Attorney) {
	x.Type = "http://schema.org/Attorney"

	return
}
