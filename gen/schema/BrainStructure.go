package schema

type BrainStructureTrait struct {

}

type BrainStructure struct {
	MetaTrait
	BrainStructureTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewBrainStructure() (x BrainStructure) {
	x.Type = "http://schema.org/BrainStructure"

	return
}
