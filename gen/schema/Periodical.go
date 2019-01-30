package schema

type PeriodicalTrait struct {

}

type Periodical struct {
	MetaTrait
	PeriodicalTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewPeriodical() (x Periodical) {
	x.Type = "http://schema.org/Periodical"

	return
}
