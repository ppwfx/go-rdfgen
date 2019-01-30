package schema

type RoofingContractorTrait struct {

}

type RoofingContractor struct {
	MetaTrait
	RoofingContractorTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewRoofingContractor() (x RoofingContractor) {
	x.Type = "http://schema.org/RoofingContractor"

	return
}
