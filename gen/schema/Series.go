package schema

type SeriesTrait struct {

}

type Series struct {
	MetaTrait
	SeriesTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSeries() (x Series) {
	x.Type = "http://schema.org/Series"

	return
}
