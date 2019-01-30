package schema

type PaintingTrait struct {

}

type Painting struct {
	MetaTrait
	PaintingTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewPainting() (x Painting) {
	x.Type = "http://schema.org/Painting"

	return
}
