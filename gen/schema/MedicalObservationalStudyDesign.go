package schema

type MedicalObservationalStudyDesignTrait struct {

}

type MedicalObservationalStudyDesign struct {
	MetaTrait
	MedicalObservationalStudyDesignTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalObservationalStudyDesign() (x MedicalObservationalStudyDesign) {
	x.Type = "http://schema.org/MedicalObservationalStudyDesign"

	return
}
