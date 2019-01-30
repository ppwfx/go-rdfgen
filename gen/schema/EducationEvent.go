package schema

type EducationEventTrait struct {

}

type EducationEvent struct {
	MetaTrait
	EducationEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewEducationEvent() (x EducationEvent) {
	x.Type = "http://schema.org/EducationEvent"

	return
}
