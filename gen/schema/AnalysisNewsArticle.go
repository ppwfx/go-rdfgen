package schema

type AnalysisNewsArticleTrait struct {

}

type AnalysisNewsArticle struct {
	MetaTrait
	AnalysisNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAnalysisNewsArticle() (x AnalysisNewsArticle) {
	x.Type = "http://schema.org/AnalysisNewsArticle"

	return
}
