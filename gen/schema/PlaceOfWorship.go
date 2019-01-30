package schema

type PlaceOfWorshipTrait struct {

}

type PlaceOfWorship struct {
	MetaTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPlaceOfWorship() (x PlaceOfWorship) {
	x.Type = "http://schema.org/PlaceOfWorship"

	return
}
