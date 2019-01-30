package schema

type BuddhistTempleTrait struct {

}

type BuddhistTemple struct {
	MetaTrait
	BuddhistTempleTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBuddhistTemple() (x BuddhistTemple) {
	x.Type = "http://schema.org/BuddhistTemple"

	return
}
