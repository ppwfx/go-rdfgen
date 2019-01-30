package schema

type MedicalGuidelineTrait struct {

	// Source of the data used to formulate the guidance, e.g. RCT, consensus opinion, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EvidenceOrigin string `json:"evidenceOrigin,omitempty" jsonld:"http://schema.org/evidenceOrigin"`

	// Date on which this guideline's recommendation was made.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	GuidelineDate *Date `json:"guidelineDate,omitempty" jsonld:"http://schema.org/guidelineDate"`

	// Strength of evidence of the data used to formulate the guideline (enumerated).
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEvidenceLevel
	//
	EvidenceLevel *MedicalEvidenceLevel `json:"evidenceLevel,omitempty" jsonld:"http://schema.org/evidenceLevel"`

	// The medical conditions, treatments, etc. that are the subject of the guideline.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	GuidelineSubject *MedicalEntity `json:"guidelineSubject,omitempty" jsonld:"http://schema.org/guidelineSubject"`

}

type MedicalGuideline struct {
	MetaTrait
	MedicalGuidelineTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalGuideline() (x MedicalGuideline) {
	x.Type = "http://schema.org/MedicalGuideline"

	return
}
