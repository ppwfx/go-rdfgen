package schema

type MuscleTrait struct {

	// Obsolete term for <a class=\"localLink\" href=\"http://schema.org/muscleAction\">muscleAction</a>. Not to be confused with <a class=\"localLink\" href=\"http://schema.org/potentialAction\">potentialAction</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Action string `json:"action,omitempty" jsonld:"http://schema.org/action"`

	// The movement the muscle generates.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MuscleAction string `json:"muscleAction,omitempty" jsonld:"http://schema.org/muscleAction"`

	// The muscle whose action counteracts the specified muscle.
	//
	// RangeIncludes:
	// - http://schema.org/Muscle
	//
	Antagonist *Muscle `json:"antagonist,omitempty" jsonld:"http://schema.org/antagonist"`

	// The blood vessel that carries blood from the heart to the muscle.
	//
	// RangeIncludes:
	// - http://schema.org/Vessel
	//
	BloodSupply *Vessel `json:"bloodSupply,omitempty" jsonld:"http://schema.org/bloodSupply"`

	// The place of attachment of a muscle, or what the muscle moves.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	Insertion *AnatomicalStructure `json:"insertion,omitempty" jsonld:"http://schema.org/insertion"`

	// The underlying innervation associated with the muscle.
	//
	// RangeIncludes:
	// - http://schema.org/Nerve
	//
	Nerve *Nerve `json:"nerve,omitempty" jsonld:"http://schema.org/nerve"`

	// The place or point where a muscle arises.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	//
	Origin *AnatomicalStructure `json:"origin,omitempty" jsonld:"http://schema.org/origin"`

}

type Muscle struct {
	MetaTrait
	MuscleTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMuscle() (x Muscle) {
	x.Type = "http://schema.org/Muscle"

	return
}
