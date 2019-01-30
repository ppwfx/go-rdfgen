package schema

type EventStatusTypeTrait struct {

}

type EventStatusType struct {
	MetaTrait
	EventStatusTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEventStatusType() (x EventStatusType) {
	x.Type = "http://schema.org/EventStatusType"

	return
}
