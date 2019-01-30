package schema

type LiteraryEventTrait struct {

}

type LiteraryEvent struct {
	MetaTrait
	LiteraryEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewLiteraryEvent() (x LiteraryEvent) {
	x.Type = "http://schema.org/LiteraryEvent"

	return
}
