package schema

type ComedyEventTrait struct {

}

type ComedyEvent struct {
	MetaTrait
	ComedyEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewComedyEvent() (x ComedyEvent) {
	x.Type = "http://schema.org/ComedyEvent"

	return
}
