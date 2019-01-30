package schema

type ProfessionalServiceTrait struct {

}

type ProfessionalService struct {
	MetaTrait
	ProfessionalServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewProfessionalService() (x ProfessionalService) {
	x.Type = "http://schema.org/ProfessionalService"

	return
}
