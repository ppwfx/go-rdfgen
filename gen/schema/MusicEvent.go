package schema

type MusicEventTrait struct {

}

type MusicEvent struct {
	MetaTrait
	MusicEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicEvent() (x MusicEvent) {
	x.Type = "http://schema.org/MusicEvent"

	return
}
