package schema

type WPFooterTrait struct {

}

type WPFooter struct {
	MetaTrait
	WPFooterTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWPFooter() (x WPFooter) {
	x.Type = "http://schema.org/WPFooter"

	return
}
