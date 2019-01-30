package schema

type SeasonTrait struct {

}

type Season struct {
	MetaTrait
	SeasonTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSeason() (x Season) {
	x.Type = "http://schema.org/Season"

	return
}
