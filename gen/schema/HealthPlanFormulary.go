package schema

type HealthPlanFormularyTrait struct {

	// Whether The costs to the patient for services under this network or formulary.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	HealthPlanCostSharing bool `json:"healthPlanCostSharing,omitempty" jsonld:"http://schema.org/healthPlanCostSharing"`

	// The tier(s) of drugs offered by this formulary or insurance plan.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanDrugTier string `json:"healthPlanDrugTier,omitempty" jsonld:"http://schema.org/healthPlanDrugTier"`

	// Whether prescriptions can be delivered by mail.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	OffersPrescriptionByMail bool `json:"offersPrescriptionByMail,omitempty" jsonld:"http://schema.org/offersPrescriptionByMail"`

}

type HealthPlanFormulary struct {
	MetaTrait
	HealthPlanFormularyTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHealthPlanFormulary() (x HealthPlanFormulary) {
	x.Type = "http://schema.org/HealthPlanFormulary"

	return
}
