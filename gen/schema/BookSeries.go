package schema

type BookSeriesTrait struct {

}

type BookSeries struct {
	MetaTrait
	BookSeriesTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewBookSeries() (x BookSeries) {
	x.Type = "http://schema.org/BookSeries"

	return
}
