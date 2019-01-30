package schema

type GeneralContractorTrait struct {

}

type GeneralContractor struct {
	MetaTrait
	GeneralContractorTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGeneralContractor() (x GeneralContractor) {
	x.Type = "http://schema.org/GeneralContractor"

	return
}
