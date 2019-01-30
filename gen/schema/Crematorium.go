package schema

type CrematoriumTrait struct {

}

type Crematorium struct {
	MetaTrait
	CrematoriumTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCrematorium() (x Crematorium) {
	x.Type = "http://schema.org/Crematorium"

	return
}
