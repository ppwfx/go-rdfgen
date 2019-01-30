package schema

type CollectionPageTrait struct {

}

type CollectionPage struct {
	MetaTrait
	CollectionPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCollectionPage() (x CollectionPage) {
	x.Type = "http://schema.org/CollectionPage"

	return
}
