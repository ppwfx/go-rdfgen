package schema

type MedicalDeviceTrait struct {

	// A description of the workup, testing, and other preparations required before implanting this device.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PreOp string `json:"preOp,omitempty" jsonld:"http://schema.org/preOp"`

	// A description of the procedure involved in setting up, using, and/or installing the device.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Procedure string `json:"procedure,omitempty" jsonld:"http://schema.org/procedure"`

	// A contraindication for this therapy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalContraindication
	//
	Contraindication interface{} `json:"contraindication,omitempty" jsonld:"http://schema.org/contraindication"`

	// A description of the postoperative procedures, care, and/or followups for this device.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PostOp string `json:"postOp,omitempty" jsonld:"http://schema.org/postOp"`

	// A goal towards an action is taken. Can be concrete or abstract.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	// - http://schema.org/MedicalDevicePurpose
	//
	Purpose interface{} `json:"purpose,omitempty" jsonld:"http://schema.org/purpose"`

	// A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalIndication
	//
	Indication *MedicalIndication `json:"indication,omitempty" jsonld:"http://schema.org/indication"`

	// A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	AdverseOutcome *MedicalEntity `json:"adverseOutcome,omitempty" jsonld:"http://schema.org/adverseOutcome"`

	// A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalEntity
	//
	SeriousAdverseOutcome *MedicalEntity `json:"seriousAdverseOutcome,omitempty" jsonld:"http://schema.org/seriousAdverseOutcome"`

}

type MedicalDevice struct {
	MetaTrait
	MedicalDeviceTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalDevice() (x MedicalDevice) {
	x.Type = "http://schema.org/MedicalDevice"

	return
}
