package schema

type FinancialServiceTrait struct {

	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification,omitempty" jsonld:"http://schema.org/feesAndCommissionsSpecification"`

}

type FinancialService struct {
	MetaTrait
	FinancialServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFinancialService() (x FinancialService) {
	x.Type = "http://schema.org/FinancialService"

	return
}
