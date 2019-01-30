package schema

type ReviewNewsArticleTrait struct {

}

type ReviewNewsArticle struct {
	MetaTrait
	ReviewNewsArticleTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	CriticReviewTrait
	ReviewTrait
	AdditionalTrait
}

func NewReviewNewsArticle() (x ReviewNewsArticle) {
	x.Type = "http://schema.org/ReviewNewsArticle"

	return
}
