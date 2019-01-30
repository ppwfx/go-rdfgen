package schema

type DrugStrengthTrait struct {

	// An active ingredient, typically chemical compounds and/or biologic substances.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ActiveIngredient string `json:"activeIngredient,omitempty" jsonld:"http://schema.org/activeIngredient"`

	// The units of an active ingredient's strength, e.g. mg.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	StrengthUnit string `json:"strengthUnit,omitempty" jsonld:"http://schema.org/strengthUnit"`

	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	//
	// RangeIncludes:
	// - http://schema.org/MaximumDoseSchedule
	//
	MaximumIntake *MaximumDoseSchedule `json:"maximumIntake,omitempty" jsonld:"http://schema.org/maximumIntake"`

	// The value of an active ingredient's strength, e.g. 325.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	StrengthValue float64 `json:"strengthValue,omitempty" jsonld:"http://schema.org/strengthValue"`

	// The location in which the strength is available.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	AvailableIn *AdministrativeArea `json:"availableIn,omitempty" jsonld:"http://schema.org/availableIn"`

}

type DrugStrength struct {
	MetaTrait
	DrugStrengthTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugStrength() (x DrugStrength) {
	x.Type = "http://schema.org/DrugStrength"

	return
}
