package schema

type ComicSeriesTrait struct {

}

type ComicSeries struct {
	MetaTrait
	ComicSeriesTrait
	PeriodicalTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewComicSeries() (x ComicSeries) {
	x.Type = "http://schema.org/ComicSeries"

	return
}
