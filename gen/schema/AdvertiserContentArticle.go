package schema

type AdvertiserContentArticleTrait struct {

}

type AdvertiserContentArticle struct {
	MetaTrait
	AdvertiserContentArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAdvertiserContentArticle() (x AdvertiserContentArticle) {
	x.Type = "http://schema.org/AdvertiserContentArticle"

	return
}
