package schema

type TheaterEventTrait struct {

}

type TheaterEvent struct {
	MetaTrait
	TheaterEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewTheaterEvent() (x TheaterEvent) {
	x.Type = "http://schema.org/TheaterEvent"

	return
}
