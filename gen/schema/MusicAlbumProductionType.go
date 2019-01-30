package schema

type MusicAlbumProductionTypeTrait struct {

}

type MusicAlbumProductionType struct {
	MetaTrait
	MusicAlbumProductionTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicAlbumProductionType() (x MusicAlbumProductionType) {
	x.Type = "http://schema.org/MusicAlbumProductionType"

	return
}
