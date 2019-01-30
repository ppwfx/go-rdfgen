package schema

type BackgroundNewsArticleTrait struct {

}

type BackgroundNewsArticle struct {
	MetaTrait
	BackgroundNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewBackgroundNewsArticle() (x BackgroundNewsArticle) {
	x.Type = "http://schema.org/BackgroundNewsArticle"

	return
}
