package schema

type MedicalCauseTrait struct {

	// The condition, complication, symptom, sign, etc. caused.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	CauseOf *MedicalEntity `json:"causeOf,omitempty" jsonld:"http://schema.org/causeOf"`

}

type MedicalCause struct {
	MetaTrait
	MedicalCauseTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalCause() (x MedicalCause) {
	x.Type = "http://schema.org/MedicalCause"

	return
}
