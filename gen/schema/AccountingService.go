package schema

type AccountingServiceTrait struct {

}

type AccountingService struct {
	MetaTrait
	AccountingServiceTrait
	FinancialServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAccountingService() (x AccountingService) {
	x.Type = "http://schema.org/AccountingService"

	return
}
