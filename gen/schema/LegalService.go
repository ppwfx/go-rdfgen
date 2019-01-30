package schema

type LegalServiceTrait struct {

}

type LegalService struct {
	MetaTrait
	LegalServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLegalService() (x LegalService) {
	x.Type = "http://schema.org/LegalService"

	return
}
