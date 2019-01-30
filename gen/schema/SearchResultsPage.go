package schema

type SearchResultsPageTrait struct {

}

type SearchResultsPage struct {
	MetaTrait
	SearchResultsPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSearchResultsPage() (x SearchResultsPage) {
	x.Type = "http://schema.org/SearchResultsPage"

	return
}
