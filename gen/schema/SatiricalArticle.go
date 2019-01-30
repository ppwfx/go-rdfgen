package schema

type SatiricalArticleTrait struct {

}

type SatiricalArticle struct {
	MetaTrait
	SatiricalArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSatiricalArticle() (x SatiricalArticle) {
	x.Type = "http://schema.org/SatiricalArticle"

	return
}
