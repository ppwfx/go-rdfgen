package schema

type NewspaperTrait struct {

}

type Newspaper struct {
	MetaTrait
	NewspaperTrait
	PeriodicalTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewNewspaper() (x Newspaper) {
	x.Type = "http://schema.org/Newspaper"

	return
}
