package schema

type MedicalProcedureTrait struct {

	// Expected or actual outcomes of the study.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalEntity
	//
	Outcome interface{} `json:"outcome,omitempty" jsonld:"http://schema.org/outcome"`

	// The status of the study (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/EventStatusType
	// - http://schema.org/MedicalStudyStatus
	//
	Status interface{} `json:"status,omitempty" jsonld:"http://schema.org/status"`

	// Typical or recommended followup care after the procedure is performed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Followup string `json:"followup,omitempty" jsonld:"http://schema.org/followup"`

	// How the procedure is performed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HowPerformed string `json:"howPerformed,omitempty" jsonld:"http://schema.org/howPerformed"`

	// Typical preparation that a patient must undergo before having the procedure performed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalEntity
	//
	Preparation interface{} `json:"preparation,omitempty" jsonld:"http://schema.org/preparation"`

	// Location in the body of the anatomical structure.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BodyLocation string `json:"bodyLocation,omitempty" jsonld:"http://schema.org/bodyLocation"`

	// A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalIndication
	//
	Indication *MedicalIndication `json:"indication,omitempty" jsonld:"http://schema.org/indication"`

	// The type of procedure, for example Surgical, Noninvasive, or Percutaneous.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalProcedureType
	//
	ProcedureType *MedicalProcedureType `json:"procedureType,omitempty" jsonld:"http://schema.org/procedureType"`

}

type MedicalProcedure struct {
	MetaTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalProcedure() (x MedicalProcedure) {
	x.Type = "http://schema.org/MedicalProcedure"

	return
}
