package schema

type MedicalBusinessTrait struct {

}

type MedicalBusiness struct {
	MetaTrait
	MedicalBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMedicalBusiness() (x MedicalBusiness) {
	x.Type = "http://schema.org/MedicalBusiness"

	return
}
