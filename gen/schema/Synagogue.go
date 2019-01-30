package schema

type SynagogueTrait struct {

}

type Synagogue struct {
	MetaTrait
	SynagogueTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewSynagogue() (x Synagogue) {
	x.Type = "http://schema.org/Synagogue"

	return
}
