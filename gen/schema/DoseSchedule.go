package schema

type DoseScheduleTrait struct {

	// Characteristics of the population for which this is intended, or which typically uses it, e.g. 'adults'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TargetPopulation string `json:"targetPopulation,omitempty" jsonld:"http://schema.org/targetPopulation"`

	// The unit of the dose, e.g. 'mg'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DoseUnit string `json:"doseUnit,omitempty" jsonld:"http://schema.org/doseUnit"`

	// How often the dose is taken, e.g. 'daily'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Frequency string `json:"frequency,omitempty" jsonld:"http://schema.org/frequency"`

	// The value of the dose, e.g. 500.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	// - http://schema.org/QualitativeValue
	//
	DoseValue interface{} `json:"doseValue,omitempty" jsonld:"http://schema.org/doseValue"`

}

type DoseSchedule struct {
	MetaTrait
	DoseScheduleTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDoseSchedule() (x DoseSchedule) {
	x.Type = "http://schema.org/DoseSchedule"

	return
}
