package schema

type ContactPageTrait struct {

}

type ContactPage struct {
	MetaTrait
	ContactPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewContactPage() (x ContactPage) {
	x.Type = "http://schema.org/ContactPage"

	return
}
