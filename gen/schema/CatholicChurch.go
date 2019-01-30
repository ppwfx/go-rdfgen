package schema

type CatholicChurchTrait struct {

}

type CatholicChurch struct {
	MetaTrait
	CatholicChurchTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCatholicChurch() (x CatholicChurch) {
	x.Type = "http://schema.org/CatholicChurch"

	return
}
