package schema

type OpinionNewsArticleTrait struct {

}

type OpinionNewsArticle struct {
	MetaTrait
	OpinionNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewOpinionNewsArticle() (x OpinionNewsArticle) {
	x.Type = "http://schema.org/OpinionNewsArticle"

	return
}
