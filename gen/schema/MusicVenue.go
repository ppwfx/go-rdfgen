package schema

type MusicVenueTrait struct {

}

type MusicVenue struct {
	MetaTrait
	MusicVenueTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicVenue() (x MusicVenue) {
	x.Type = "http://schema.org/MusicVenue"

	return
}
