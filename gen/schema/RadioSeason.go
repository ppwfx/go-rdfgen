package schema

type RadioSeasonTrait struct {

}

type RadioSeason struct {
	MetaTrait
	RadioSeasonTrait
	CreativeWorkSeasonTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewRadioSeason() (x RadioSeason) {
	x.Type = "http://schema.org/RadioSeason"

	return
}
