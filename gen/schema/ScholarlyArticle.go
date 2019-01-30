package schema

type ScholarlyArticleTrait struct {

}

type ScholarlyArticle struct {
	MetaTrait
	ScholarlyArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewScholarlyArticle() (x ScholarlyArticle) {
	x.Type = "http://schema.org/ScholarlyArticle"

	return
}
