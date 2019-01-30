package schema

type ConsortiumTrait struct {

}

type Consortium struct {
	MetaTrait
	ConsortiumTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewConsortium() (x Consortium) {
	x.Type = "http://schema.org/Consortium"

	return
}
