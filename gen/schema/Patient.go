package schema

type PatientTrait struct {

	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	HealthCondition *MedicalCondition `json:"healthCondition,omitempty" jsonld:"http://schema.org/healthCondition"`

	// One or more alternative conditions considered in the differential diagnosis process as output of a diagnosis process.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	Diagnosis *MedicalCondition `json:"diagnosis,omitempty" jsonld:"http://schema.org/diagnosis"`

	// Specifying a drug or medicine used in a medication procedure
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	Drug *Drug `json:"drug,omitempty" jsonld:"http://schema.org/drug"`

}

type Patient struct {
	MetaTrait
	PatientTrait
	PersonTrait
	ThingTrait
	MedicalAudienceTrait
	AudienceTrait
	IntangibleTrait
	MedicalEnumerationTrait
	EnumerationTrait
	PeopleAudienceTrait
	AdditionalTrait
}

func NewPatient() (x Patient) {
	x.Type = "http://schema.org/Patient"

	return
}
