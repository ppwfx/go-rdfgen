package schema

type InsuranceAgencyTrait struct {

}

type InsuranceAgency struct {
	MetaTrait
	InsuranceAgencyTrait
	FinancialServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewInsuranceAgency() (x InsuranceAgency) {
	x.Type = "http://schema.org/InsuranceAgency"

	return
}
