package schema

type PharmacyTrait struct {

}

type Pharmacy struct {
	MetaTrait
	PharmacyTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	MedicalBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	AdditionalTrait
}

func NewPharmacy() (x Pharmacy) {
	x.Type = "http://schema.org/Pharmacy"

	return
}
