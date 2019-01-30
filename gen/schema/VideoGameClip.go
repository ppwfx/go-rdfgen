package schema

type VideoGameClipTrait struct {

}

type VideoGameClip struct {
	MetaTrait
	VideoGameClipTrait
	ClipTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewVideoGameClip() (x VideoGameClip) {
	x.Type = "http://schema.org/VideoGameClip"

	return
}
