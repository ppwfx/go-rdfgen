package schema

type DrugCostTrait struct {

	// The cost per unit of the drug.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	// - http://schema.org/QualitativeValue
	//
	CostPerUnit interface{} `json:"costPerUnit,omitempty" jsonld:"http://schema.org/costPerUnit"`

	// The unit in which the drug is measured, e.g. '5 mg tablet'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DrugUnit string `json:"drugUnit,omitempty" jsonld:"http://schema.org/drugUnit"`

	// The currency (in 3-letter of the drug cost. See: http://en.wikipedia.org/wiki/ISO_4217
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CostCurrency string `json:"costCurrency,omitempty" jsonld:"http://schema.org/costCurrency"`

	// Additional details to capture the origin of the cost data. For example, 'Medicare Part B'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CostOrigin string `json:"costOrigin,omitempty" jsonld:"http://schema.org/costOrigin"`

	// The location in which the status applies.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	ApplicableLocation *AdministrativeArea `json:"applicableLocation,omitempty" jsonld:"http://schema.org/applicableLocation"`

	// The category of cost, such as wholesale, retail, reimbursement cap, etc.
	//
	// RangeIncludes:
	// - http://schema.org/DrugCostCategory
	//
	CostCategory *DrugCostCategory `json:"costCategory,omitempty" jsonld:"http://schema.org/costCategory"`

}

type DrugCost struct {
	MetaTrait
	DrugCostTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugCost() (x DrugCost) {
	x.Type = "http://schema.org/DrugCost"

	return
}
