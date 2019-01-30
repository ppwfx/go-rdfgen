package schema

type WPSideBarTrait struct {

}

type WPSideBar struct {
	MetaTrait
	WPSideBarTrait
	WebPageElementTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWPSideBar() (x WPSideBar) {
	x.Type = "http://schema.org/WPSideBar"

	return
}
