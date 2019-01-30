package schema

type MedicalEvidenceLevelTrait struct {

}

type MedicalEvidenceLevel struct {
	MetaTrait
	MedicalEvidenceLevelTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalEvidenceLevel() (x MedicalEvidenceLevel) {
	x.Type = "http://schema.org/MedicalEvidenceLevel"

	return
}
