package schema

type MusicStoreTrait struct {

}

type MusicStore struct {
	MetaTrait
	MusicStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMusicStore() (x MusicStore) {
	x.Type = "http://schema.org/MusicStore"

	return
}
