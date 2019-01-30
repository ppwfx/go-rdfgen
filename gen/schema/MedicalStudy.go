package schema

type MedicalStudyTrait struct {

	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Sponsor interface{} `json:"sponsor,omitempty" jsonld:"http://schema.org/sponsor"`

	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	HealthCondition *MedicalCondition `json:"healthCondition,omitempty" jsonld:"http://schema.org/healthCondition"`

	// Expected or actual outcomes of the study.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalEntity
	//
	Outcome interface{} `json:"outcome,omitempty" jsonld:"http://schema.org/outcome"`

	// Any characteristics of the population used in the study, e.g. 'males under 65'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Population string `json:"population,omitempty" jsonld:"http://schema.org/population"`

	// The status of the study (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/EventStatusType
	// - http://schema.org/MedicalStudyStatus
	//
	Status interface{} `json:"status,omitempty" jsonld:"http://schema.org/status"`

	// The location in which the study is taking/took place.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	StudyLocation *AdministrativeArea `json:"studyLocation,omitempty" jsonld:"http://schema.org/studyLocation"`

	// A subject of the study, i.e. one of the medical conditions, therapies, devices, drugs, etc. investigated by the study.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	StudySubject *MedicalEntity `json:"studySubject,omitempty" jsonld:"http://schema.org/studySubject"`

}

type MedicalStudy struct {
	MetaTrait
	MedicalStudyTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalStudy() (x MedicalStudy) {
	x.Type = "http://schema.org/MedicalStudy"

	return
}
