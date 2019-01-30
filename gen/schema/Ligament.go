package schema

type LigamentTrait struct {

}

type Ligament struct {
	MetaTrait
	LigamentTrait
	AnatomicalStructureTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewLigament() (x Ligament) {
	x.Type = "http://schema.org/Ligament"

	return
}
