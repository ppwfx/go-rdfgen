package schema

type WPAdBlockTrait struct {

}

type WPAdBlock struct {
	MetaTrait
	WPAdBlockTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWPAdBlock() (x WPAdBlock) {
	x.Type = "http://schema.org/WPAdBlock"

	return
}
