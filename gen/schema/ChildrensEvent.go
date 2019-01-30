package schema

type ChildrensEventTrait struct {

}

type ChildrensEvent struct {
	MetaTrait
	ChildrensEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewChildrensEvent() (x ChildrensEvent) {
	x.Type = "http://schema.org/ChildrensEvent"

	return
}
