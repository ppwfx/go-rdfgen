package schema

type DietarySupplementTrait struct {

	// The manufacturer of the product.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Manufacturer *Organization `json:"manufacturer,omitempty" jsonld:"http://schema.org/manufacturer"`

	// The drug or supplement's legal status, including any controlled substance schedules that apply.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/DrugLegalStatus
	// - http://schema.org/MedicalEnumeration
	//
	LegalStatus interface{} `json:"legalStatus,omitempty" jsonld:"http://schema.org/legalStatus"`

	// Any potential safety concern associated with the supplement. May include interactions with other drugs and foods, pregnancy, breastfeeding, known adverse reactions, and documented efficacy of the supplement.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SafetyConsideration string `json:"safetyConsideration,omitempty" jsonld:"http://schema.org/safetyConsideration"`

	// Characteristics of the population for which this is intended, or which typically uses it, e.g. 'adults'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TargetPopulation string `json:"targetPopulation,omitempty" jsonld:"http://schema.org/targetPopulation"`

	// An active ingredient, typically chemical compounds and/or biologic substances.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ActiveIngredient string `json:"activeIngredient,omitempty" jsonld:"http://schema.org/activeIngredient"`

	// The generic name of this drug or supplement.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	NonProprietaryName string `json:"nonProprietaryName,omitempty" jsonld:"http://schema.org/nonProprietaryName"`

	// The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MechanismOfAction string `json:"mechanismOfAction,omitempty" jsonld:"http://schema.org/mechanismOfAction"`

	// Descriptive information establishing a historical perspective on the supplement. May include the rationale for the name, the population where the supplement first came to prominence, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Background string `json:"background,omitempty" jsonld:"http://schema.org/background"`

	// Proprietary name given to the diet plan, typically by its originator or creator.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProprietaryName string `json:"proprietaryName,omitempty" jsonld:"http://schema.org/proprietaryName"`

	// True if this item's name is a proprietary/brand name (vs. generic name).
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsProprietary bool `json:"isProprietary,omitempty" jsonld:"http://schema.org/isProprietary"`

	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	//
	// RangeIncludes:
	// - http://schema.org/MaximumDoseSchedule
	//
	MaximumIntake *MaximumDoseSchedule `json:"maximumIntake,omitempty" jsonld:"http://schema.org/maximumIntake"`

	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	//
	// RangeIncludes:
	// - http://schema.org/RecommendedDoseSchedule
	//
	RecommendedIntake *RecommendedDoseSchedule `json:"recommendedIntake,omitempty" jsonld:"http://schema.org/recommendedIntake"`

}

type DietarySupplement struct {
	MetaTrait
	DietarySupplementTrait
	SubstanceTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDietarySupplement() (x DietarySupplement) {
	x.Type = "http://schema.org/DietarySupplement"

	return
}
