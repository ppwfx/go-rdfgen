package schema

type AskPublicNewsArticleTrait struct {

}

type AskPublicNewsArticle struct {
	MetaTrait
	AskPublicNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAskPublicNewsArticle() (x AskPublicNewsArticle) {
	x.Type = "http://schema.org/AskPublicNewsArticle"

	return
}
