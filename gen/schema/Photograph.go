package schema

type PhotographTrait struct {

}

type Photograph struct {
	MetaTrait
	PhotographTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewPhotograph() (x Photograph) {
	x.Type = "http://schema.org/Photograph"

	return
}
