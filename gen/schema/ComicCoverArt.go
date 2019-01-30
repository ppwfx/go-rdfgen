package schema

type ComicCoverArtTrait struct {

}

type ComicCoverArt struct {
	MetaTrait
	ComicCoverArtTrait
	ComicStoryTrait
	CreativeWorkTrait
	ThingTrait
	CoverArtTrait
	VisualArtworkTrait
	AdditionalTrait
}

func NewComicCoverArt() (x ComicCoverArt) {
	x.Type = "http://schema.org/ComicCoverArt"

	return
}
