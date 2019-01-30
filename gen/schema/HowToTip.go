package schema

type HowToTipTrait struct {

}

type HowToTip struct {
	MetaTrait
	HowToTipTrait
	CreativeWorkTrait
	ThingTrait
	ListItemTrait
	IntangibleTrait
	AdditionalTrait
}

func NewHowToTip() (x HowToTip) {
	x.Type = "http://schema.org/HowToTip"

	return
}
