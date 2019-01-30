package schema

type RadioEpisodeTrait struct {

}

type RadioEpisode struct {
	MetaTrait
	RadioEpisodeTrait
	EpisodeTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewRadioEpisode() (x RadioEpisode) {
	x.Type = "http://schema.org/RadioEpisode"

	return
}
