package schema

type MosqueTrait struct {

}

type Mosque struct {
	MetaTrait
	MosqueTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewMosque() (x Mosque) {
	x.Type = "http://schema.org/Mosque"

	return
}
