package schema

type FestivalTrait struct {

}

type Festival struct {
	MetaTrait
	FestivalTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewFestival() (x Festival) {
	x.Type = "http://schema.org/Festival"

	return
}
