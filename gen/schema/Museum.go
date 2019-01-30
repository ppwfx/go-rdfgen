package schema

type MuseumTrait struct {

}

type Museum struct {
	MetaTrait
	MuseumTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewMuseum() (x Museum) {
	x.Type = "http://schema.org/Museum"

	return
}
