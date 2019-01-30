package schema

type RealEstateAgentTrait struct {

}

type RealEstateAgent struct {
	MetaTrait
	RealEstateAgentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewRealEstateAgent() (x RealEstateAgent) {
	x.Type = "http://schema.org/RealEstateAgent"

	return
}
