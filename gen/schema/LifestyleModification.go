package schema

type LifestyleModificationTrait struct {

}

type LifestyleModification struct {
	MetaTrait
	LifestyleModificationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewLifestyleModification() (x LifestyleModification) {
	x.Type = "http://schema.org/LifestyleModification"

	return
}
