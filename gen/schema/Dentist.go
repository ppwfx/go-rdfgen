package schema

type DentistTrait struct {

}

type Dentist struct {
	MetaTrait
	DentistTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	MedicalOrganizationTrait
	MedicalBusinessTrait
	AdditionalTrait
}

func NewDentist() (x Dentist) {
	x.Type = "http://schema.org/Dentist"

	return
}
