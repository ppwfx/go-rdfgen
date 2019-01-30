package schema

type MedicalCodeTrait struct {

	// The coding system, e.g. 'ICD-10'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CodingSystem string `json:"codingSystem,omitempty" jsonld:"http://schema.org/codingSystem"`

	// A short textual code that uniquely identifies the value.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CodeValue string `json:"codeValue,omitempty" jsonld:"http://schema.org/codeValue"`

}

type MedicalCode struct {
	MetaTrait
	MedicalCodeTrait
	CategoryCodeTrait
	DefinedTermTrait
	IntangibleTrait
	ThingTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	AdditionalTrait
}

func NewMedicalCode() (x MedicalCode) {
	x.Type = "http://schema.org/MedicalCode"

	return
}
