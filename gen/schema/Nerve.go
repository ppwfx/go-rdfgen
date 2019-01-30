package schema

type NerveTrait struct {

	// The branches that delineate from the nerve bundle. Not to be confused with <a class=\"localLink\" href=\"http://schema.org/branchOf\">branchOf</a>.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	Branch *AnatomicalStructure `json:"branch,omitempty" jsonld:"http://schema.org/branch"`

	// The neurological pathway extension that involves muscle control.
	//
	// RangeIncludes:
	// - http://schema.org/Muscle
	//
	NerveMotor *Muscle `json:"nerveMotor,omitempty" jsonld:"http://schema.org/nerveMotor"`

	// The neurological pathway extension that inputs and sends information to the brain or spinal cord.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/SuperficialAnatomy
	//
	SensoryUnit interface{} `json:"sensoryUnit,omitempty" jsonld:"http://schema.org/sensoryUnit"`

	// The neurological pathway that originates the neurons.
	//
	// RangeIncludes:
	// - http://schema.org/BrainStructure
	//
	SourcedFrom *BrainStructure `json:"sourcedFrom,omitempty" jsonld:"http://schema.org/sourcedFrom"`

}

type Nerve struct {
	MetaTrait
	NerveTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewNerve() (x Nerve) {
	x.Type = "http://schema.org/Nerve"

	return
}
