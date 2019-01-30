package schema

type ReportageNewsArticleTrait struct {

}

type ReportageNewsArticle struct {
	MetaTrait
	ReportageNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewReportageNewsArticle() (x ReportageNewsArticle) {
	x.Type = "http://schema.org/ReportageNewsArticle"

	return
}
