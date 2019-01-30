package schema

type CoverArtTrait struct {

}

type CoverArt struct {
	MetaTrait
	CoverArtTrait
	VisualArtworkTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCoverArt() (x CoverArt) {
	x.Type = "http://schema.org/CoverArt"

	return
}
