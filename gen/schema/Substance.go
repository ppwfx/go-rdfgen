package schema

type SubstanceTrait struct {

	// An active ingredient, typically chemical compounds and/or biologic substances.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ActiveIngredient string `json:"activeIngredient,omitempty" jsonld:"http://schema.org/activeIngredient"`

	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	//
	// RangeIncludes:
	// - http://schema.org/MaximumDoseSchedule
	//
	MaximumIntake *MaximumDoseSchedule `json:"maximumIntake,omitempty" jsonld:"http://schema.org/maximumIntake"`

}

type Substance struct {
	MetaTrait
	SubstanceTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewSubstance() (x Substance) {
	x.Type = "http://schema.org/Substance"

	return
}
