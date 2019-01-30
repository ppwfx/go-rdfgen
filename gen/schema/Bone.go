package schema

type BoneTrait struct {

}

type Bone struct {
	MetaTrait
	BoneTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewBone() (x Bone) {
	x.Type = "http://schema.org/Bone"

	return
}
