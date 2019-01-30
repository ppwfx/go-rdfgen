package schema

type HealthInsurancePlanTrait struct {

	// A contact point for a person or organization.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	//
	ContactPoint *ContactPoint `json:"contactPoint,omitempty" jsonld:"http://schema.org/contactPoint"`

	// The tier(s) of drugs offered by this formulary or insurance plan.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanDrugTier string `json:"healthPlanDrugTier,omitempty" jsonld:"http://schema.org/healthPlanDrugTier"`

	// The URL that goes directly to the summary of benefits and coverage for the specific standard plan or plan variation.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	BenefitsSummaryUrl string `json:"benefitsSummaryUrl,omitempty" jsonld:"http://schema.org/benefitsSummaryUrl"`

	// TODO.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanDrugOption string `json:"healthPlanDrugOption,omitempty" jsonld:"http://schema.org/healthPlanDrugOption"`

	// The 14-character, HIOS-generated Plan ID number. (Plan IDs must be unique, even across different markets.)
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanId string `json:"healthPlanId,omitempty" jsonld:"http://schema.org/healthPlanId"`

	// The URL that goes directly to the plan brochure for the specific standard plan or plan variation.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	HealthPlanMarketingUrl string `json:"healthPlanMarketingUrl,omitempty" jsonld:"http://schema.org/healthPlanMarketingUrl"`

	// Formularies covered by this plan.
	//
	// RangeIncludes:
	// - http://schema.org/HealthPlanFormulary
	//
	IncludesHealthPlanFormulary *HealthPlanFormulary `json:"includesHealthPlanFormulary,omitempty" jsonld:"http://schema.org/includesHealthPlanFormulary"`

	// Networks covered by this plan.
	//
	// RangeIncludes:
	// - http://schema.org/HealthPlanNetwork
	//
	IncludesHealthPlanNetwork *HealthPlanNetwork `json:"includesHealthPlanNetwork,omitempty" jsonld:"http://schema.org/includesHealthPlanNetwork"`

	// The standard for interpreting thePlan ID. The preferred is \"HIOS\". See the Centers for Medicare &amp; Medicaid Services for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	UsesHealthPlanIdStandard interface{} `json:"usesHealthPlanIdStandard,omitempty" jsonld:"http://schema.org/usesHealthPlanIdStandard"`

}

type HealthInsurancePlan struct {
	MetaTrait
	HealthInsurancePlanTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHealthInsurancePlan() (x HealthInsurancePlan) {
	x.Type = "http://schema.org/HealthInsurancePlan"

	return
}
