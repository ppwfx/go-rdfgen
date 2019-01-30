package schema

type ZooTrait struct {

}

type Zoo struct {
	MetaTrait
	ZooTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewZoo() (x Zoo) {
	x.Type = "http://schema.org/Zoo"

	return
}
