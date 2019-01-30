package schema

type CemeteryTrait struct {

}

type Cemetery struct {
	MetaTrait
	CemeteryTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCemetery() (x Cemetery) {
	x.Type = "http://schema.org/Cemetery"

	return
}
