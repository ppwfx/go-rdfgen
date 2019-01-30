package schema

type SculptureTrait struct {

}

type Sculpture struct {
	MetaTrait
	SculptureTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSculpture() (x Sculpture) {
	x.Type = "http://schema.org/Sculpture"

	return
}
