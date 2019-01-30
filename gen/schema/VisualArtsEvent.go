package schema

type VisualArtsEventTrait struct {

}

type VisualArtsEvent struct {
	MetaTrait
	VisualArtsEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewVisualArtsEvent() (x VisualArtsEvent) {
	x.Type = "http://schema.org/VisualArtsEvent"

	return
}
