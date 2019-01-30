package schema

type RadioClipTrait struct {

}

type RadioClip struct {
	MetaTrait
	RadioClipTrait
	ClipTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewRadioClip() (x RadioClip) {
	x.Type = "http://schema.org/RadioClip"

	return
}
