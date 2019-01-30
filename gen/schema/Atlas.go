package schema

type AtlasTrait struct {

}

type Atlas struct {
	MetaTrait
	AtlasTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAtlas() (x Atlas) {
	x.Type = "http://schema.org/Atlas"

	return
}
