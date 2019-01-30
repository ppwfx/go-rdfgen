package schema

type MusicAlbumReleaseTypeTrait struct {

}

type MusicAlbumReleaseType struct {
	MetaTrait
	MusicAlbumReleaseTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicAlbumReleaseType() (x MusicAlbumReleaseType) {
	x.Type = "http://schema.org/MusicAlbumReleaseType"

	return
}
