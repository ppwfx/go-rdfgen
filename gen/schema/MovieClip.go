package schema

type MovieClipTrait struct {

}

type MovieClip struct {
	MetaTrait
	MovieClipTrait
	ClipTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMovieClip() (x MovieClip) {
	x.Type = "http://schema.org/MovieClip"

	return
}
