package schema

type NGOTrait struct {

}

type NGO struct {
	MetaTrait
	NGOTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewNGO() (x NGO) {
	x.Type = "http://schema.org/NGO"

	return
}
