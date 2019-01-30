package schema

type OpticianTrait struct {

}

type Optician struct {
	MetaTrait
	OpticianTrait
	MedicalBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewOptician() (x Optician) {
	x.Type = "http://schema.org/Optician"

	return
}
