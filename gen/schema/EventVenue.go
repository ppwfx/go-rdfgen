package schema

type EventVenueTrait struct {

}

type EventVenue struct {
	MetaTrait
	EventVenueTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewEventVenue() (x EventVenue) {
	x.Type = "http://schema.org/EventVenue"

	return
}
