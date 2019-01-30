package schema

type OnDemandEventTrait struct {

}

type OnDemandEvent struct {
	MetaTrait
	OnDemandEventTrait
	PublicationEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewOnDemandEvent() (x OnDemandEvent) {
	x.Type = "http://schema.org/OnDemandEvent"

	return
}
