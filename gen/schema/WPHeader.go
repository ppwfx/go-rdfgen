package schema

type WPHeaderTrait struct {

}

type WPHeader struct {
	MetaTrait
	WPHeaderTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWPHeader() (x WPHeader) {
	x.Type = "http://schema.org/WPHeader"

	return
}
