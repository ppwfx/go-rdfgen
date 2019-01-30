package schema

type AutomatedTellerTrait struct {

}

type AutomatedTeller struct {
	MetaTrait
	AutomatedTellerTrait
	FinancialServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutomatedTeller() (x AutomatedTeller) {
	x.Type = "http://schema.org/AutomatedTeller"

	return
}
