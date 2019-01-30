package schema

type NotaryTrait struct {

}

type Notary struct {
	MetaTrait
	NotaryTrait
	LegalServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewNotary() (x Notary) {
	x.Type = "http://schema.org/Notary"

	return
}
