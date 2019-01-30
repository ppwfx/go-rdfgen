package schema

type MedicalConditionTrait struct {

	// The status of the study (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/EventStatusType
	// - http://schema.org/MedicalStudyStatus
	//
	Status interface{} `json:"status,omitempty" jsonld:"http://schema.org/status"`

	// A more specific type of the condition, where applicable, for example 'Type 1 Diabetes', 'Type 2 Diabetes', or 'Gestational Diabetes' for Diabetes.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Subtype string `json:"subtype,omitempty" jsonld:"http://schema.org/subtype"`

	// The expected progression of the condition if it is not treated and allowed to progress naturally.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	NaturalProgression string `json:"naturalProgression,omitempty" jsonld:"http://schema.org/naturalProgression"`

	// Changes in the normal mechanical, physical, and biochemical functions that are associated with this activity or condition.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Pathophysiology string `json:"pathophysiology,omitempty" jsonld:"http://schema.org/pathophysiology"`

	// A possible unexpected and unfavorable evolution of a medical condition. Complications may include worsening of the signs or symptoms of the disease, extension of the condition to other organ systems, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PossibleComplication string `json:"possibleComplication,omitempty" jsonld:"http://schema.org/possibleComplication"`

	// The likely outcome in either the short term or long term of the medical condition.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExpectedPrognosis string `json:"expectedPrognosis,omitempty" jsonld:"http://schema.org/expectedPrognosis"`

	// The characteristics of associated patients, such as age, gender, race etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Epidemiology string `json:"epidemiology,omitempty" jsonld:"http://schema.org/epidemiology"`

	// The anatomy of the underlying organ system or structures associated with this entity.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	// - http://schema.org/SuperficialAnatomy
	//
	AssociatedAnatomy interface{} `json:"associatedAnatomy,omitempty" jsonld:"http://schema.org/associatedAnatomy"`

	// Specifying a drug or medicine used in a medication procedure
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	Drug *Drug `json:"drug,omitempty" jsonld:"http://schema.org/drug"`

	// Specifying a cause of something in general. e.g in medicine , one of the causative agent(s) that are most directly responsible for the pathophysiologic process that eventually results in the occurrence.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCause
	//
	Cause *MedicalCause `json:"cause,omitempty" jsonld:"http://schema.org/cause"`

	// One of a set of differential diagnoses for the condition. Specifically, a closely-related or competing diagnosis typically considered later in the cognitive process whereby this medical condition is distinguished from others most likely responsible for a similar collection of signs and symptoms to reach the most parsimonious diagnosis or diagnoses in a patient.
	//
	// RangeIncludes:
	// - http://schema.org/DDxElement
	//
	DifferentialDiagnosis *DDxElement `json:"differentialDiagnosis,omitempty" jsonld:"http://schema.org/differentialDiagnosis"`

	// A possible treatment to address this condition, sign or symptom.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	PossibleTreatment *MedicalTherapy `json:"possibleTreatment,omitempty" jsonld:"http://schema.org/possibleTreatment"`

	// A preventative therapy used to prevent an initial occurrence of the medical condition, such as vaccination.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	PrimaryPrevention *MedicalTherapy `json:"primaryPrevention,omitempty" jsonld:"http://schema.org/primaryPrevention"`

	// A modifiable or non-modifiable factor that increases the risk of a patient contracting this condition, e.g. age,  coexisting condition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalRiskFactor
	//
	RiskFactor *MedicalRiskFactor `json:"riskFactor,omitempty" jsonld:"http://schema.org/riskFactor"`

	// A preventative therapy used to prevent reoccurrence of the medical condition after an initial episode of the condition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	SecondaryPrevention *MedicalTherapy `json:"secondaryPrevention,omitempty" jsonld:"http://schema.org/secondaryPrevention"`

	// A sign or symptom of this condition. Signs are objective or physically observable manifestations of the medical condition while symptoms are the subjective experience of the medical condition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSignOrSymptom
	//
	SignOrSymptom *MedicalSignOrSymptom `json:"signOrSymptom,omitempty" jsonld:"http://schema.org/signOrSymptom"`

	// The stage of the condition, if applicable.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalConditionStage
	//
	Stage *MedicalConditionStage `json:"stage,omitempty" jsonld:"http://schema.org/stage"`

	// A medical test typically performed given this condition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTest
	//
	TypicalTest *MedicalTest `json:"typicalTest,omitempty" jsonld:"http://schema.org/typicalTest"`

}

type MedicalCondition struct {
	MetaTrait
	MedicalConditionTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalCondition() (x MedicalCondition) {
	x.Type = "http://schema.org/MedicalCondition"

	return
}
