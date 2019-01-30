package schema

type HowToToolTrait struct {

}

type HowToTool struct {
	MetaTrait
	HowToToolTrait
	HowToItemTrait
	ListItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHowToTool() (x HowToTool) {
	x.Type = "http://schema.org/HowToTool"

	return
}
