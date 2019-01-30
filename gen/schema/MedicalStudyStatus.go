package schema

type MedicalStudyStatusTrait struct {

}

type MedicalStudyStatus struct {
	MetaTrait
	MedicalStudyStatusTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalStudyStatus() (x MedicalStudyStatus) {
	x.Type = "http://schema.org/MedicalStudyStatus"

	return
}
