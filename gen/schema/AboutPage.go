package schema

type AboutPageTrait struct {

}

type AboutPage struct {
	MetaTrait
	AboutPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAboutPage() (x AboutPage) {
	x.Type = "http://schema.org/AboutPage"

	return
}
