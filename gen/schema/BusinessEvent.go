package schema

type BusinessEventTrait struct {

}

type BusinessEvent struct {
	MetaTrait
	BusinessEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewBusinessEvent() (x BusinessEvent) {
	x.Type = "http://schema.org/BusinessEvent"

	return
}
