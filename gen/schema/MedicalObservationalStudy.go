package schema

type MedicalObservationalStudyTrait struct {

	// Specifics about the observational study design (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/MedicalObservationalStudyDesign
	//
	StudyDesign *MedicalObservationalStudyDesign `json:"studyDesign,omitempty" jsonld:"http://schema.org/studyDesign"`

}

type MedicalObservationalStudy struct {
	MetaTrait
	MedicalObservationalStudyTrait
	MedicalStudyTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalObservationalStudy() (x MedicalObservationalStudy) {
	x.Type = "http://schema.org/MedicalObservationalStudy"

	return
}
