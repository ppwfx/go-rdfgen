package schema

type BusinessAudienceTrait struct {

	// The number of employees in an organization e.g. business.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees,omitempty" jsonld:"http://schema.org/numberOfEmployees"`

	// The size of the business in annual revenue.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	YearlyRevenue *QuantitativeValue `json:"yearlyRevenue,omitempty" jsonld:"http://schema.org/yearlyRevenue"`

	// The age of the business.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	YearsInOperation *QuantitativeValue `json:"yearsInOperation,omitempty" jsonld:"http://schema.org/yearsInOperation"`

}

type BusinessAudience struct {
	MetaTrait
	BusinessAudienceTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBusinessAudience() (x BusinessAudience) {
	x.Type = "http://schema.org/BusinessAudience"

	return
}
