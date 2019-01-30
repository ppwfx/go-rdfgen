package schema

type FAQPageTrait struct {

}

type FAQPage struct {
	MetaTrait
	FAQPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewFAQPage() (x FAQPage) {
	x.Type = "http://schema.org/FAQPage"

	return
}
