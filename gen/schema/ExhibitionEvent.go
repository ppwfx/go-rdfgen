package schema

type ExhibitionEventTrait struct {

}

type ExhibitionEvent struct {
	MetaTrait
	ExhibitionEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewExhibitionEvent() (x ExhibitionEvent) {
	x.Type = "http://schema.org/ExhibitionEvent"

	return
}
