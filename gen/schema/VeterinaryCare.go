package schema

type VeterinaryCareTrait struct {

}

type VeterinaryCare struct {
	MetaTrait
	VeterinaryCareTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewVeterinaryCare() (x VeterinaryCare) {
	x.Type = "http://schema.org/VeterinaryCare"

	return
}
