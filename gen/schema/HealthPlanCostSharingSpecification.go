package schema

type HealthPlanCostSharingSpecificationTrait struct {

	// Whether the coinsurance applies before or after deductible, etc. TODO: Is this a closed set?
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanCoinsuranceOption string `json:"healthPlanCoinsuranceOption,omitempty" jsonld:"http://schema.org/healthPlanCoinsuranceOption"`

	// Whether The rate of coinsurance expressed as a number between 0.0 and 1.0.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	HealthPlanCoinsuranceRate float64 `json:"healthPlanCoinsuranceRate,omitempty" jsonld:"http://schema.org/healthPlanCoinsuranceRate"`

	// Whether The copay amount.
	//
	// RangeIncludes:
	// - http://schema.org/PriceSpecification
	//
	HealthPlanCopay *PriceSpecification `json:"healthPlanCopay,omitempty" jsonld:"http://schema.org/healthPlanCopay"`

	// Whether the copay is before or after deductible, etc. TODO: Is this a closed set?
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanCopayOption string `json:"healthPlanCopayOption,omitempty" jsonld:"http://schema.org/healthPlanCopayOption"`

	// The category or type of pharmacy associated with this cost sharing.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanPharmacyCategory string `json:"healthPlanPharmacyCategory,omitempty" jsonld:"http://schema.org/healthPlanPharmacyCategory"`

}

type HealthPlanCostSharingSpecification struct {
	MetaTrait
	HealthPlanCostSharingSpecificationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHealthPlanCostSharingSpecification() (x HealthPlanCostSharingSpecification) {
	x.Type = "http://schema.org/HealthPlanCostSharingSpecification"

	return
}
