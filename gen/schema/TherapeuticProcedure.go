package schema

type TherapeuticProcedureTrait struct {

	// Specifying a drug or medicine used in a medication procedure
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	Drug *Drug `json:"drug,omitempty" jsonld:"http://schema.org/drug"`

	// A dosing schedule for the drug for a given population, either observed, recommended, or maximum dose based on the type used.
	//
	// RangeIncludes:
	// - http://schema.org/DoseSchedule
	//
	DoseSchedule *DoseSchedule `json:"doseSchedule,omitempty" jsonld:"http://schema.org/doseSchedule"`

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

}

type TherapeuticProcedure struct {
	MetaTrait
	TherapeuticProcedureTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewTherapeuticProcedure() (x TherapeuticProcedure) {
	x.Type = "http://schema.org/TherapeuticProcedure"

	return
}
