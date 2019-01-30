package schema

type MusicVideoObjectTrait struct {

}

type MusicVideoObject struct {
	MetaTrait
	MusicVideoObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicVideoObject() (x MusicVideoObject) {
	x.Type = "http://schema.org/MusicVideoObject"

	return
}
