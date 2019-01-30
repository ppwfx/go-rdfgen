package schema

type SiteNavigationElementTrait struct {

}

type SiteNavigationElement struct {
	MetaTrait
	SiteNavigationElementTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSiteNavigationElement() (x SiteNavigationElement) {
	x.Type = "http://schema.org/SiteNavigationElement"

	return
}
