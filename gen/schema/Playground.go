package schema

type PlaygroundTrait struct {

}

type Playground struct {
	MetaTrait
	PlaygroundTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPlayground() (x Playground) {
	x.Type = "http://schema.org/Playground"

	return
}
