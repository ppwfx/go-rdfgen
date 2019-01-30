package schema

type DanceEventTrait struct {

}

type DanceEvent struct {
	MetaTrait
	DanceEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewDanceEvent() (x DanceEvent) {
	x.Type = "http://schema.org/DanceEvent"

	return
}
