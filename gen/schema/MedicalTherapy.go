package schema

type MedicalTherapyTrait struct {

	// A contraindication for this therapy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalContraindication
	//
	Contraindication interface{} `json:"contraindication,omitempty" jsonld:"http://schema.org/contraindication"`

	// A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	SeriousAdverseOutcome *MedicalEntity `json:"seriousAdverseOutcome,omitempty" jsonld:"http://schema.org/seriousAdverseOutcome"`

	// A therapy that duplicates or overlaps this one.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalTherapy
	//
	DuplicateTherapy *MedicalTherapy `json:"duplicateTherapy,omitempty" jsonld:"http://schema.org/duplicateTherapy"`

}

type MedicalTherapy struct {
	MetaTrait
	MedicalTherapyTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalTherapy() (x MedicalTherapy) {
	x.Type = "http://schema.org/MedicalTherapy"

	return
}
