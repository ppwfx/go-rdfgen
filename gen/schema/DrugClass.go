package schema

type DrugClassTrait struct {

	// Specifying a drug or medicine used in a medication procedure
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	Drug *Drug `json:"drug,omitempty" jsonld:"http://schema.org/drug"`

}

type DrugClass struct {
	MetaTrait
	DrugClassTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugClass() (x DrugClass) {
	x.Type = "http://schema.org/DrugClass"

	return
}
