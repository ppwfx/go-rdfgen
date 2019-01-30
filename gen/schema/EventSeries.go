package schema

type EventSeriesTrait struct {

}

type EventSeries struct {
	MetaTrait
	EventSeriesTrait
	EventTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewEventSeries() (x EventSeries) {
	x.Type = "http://schema.org/EventSeries"

	return
}
