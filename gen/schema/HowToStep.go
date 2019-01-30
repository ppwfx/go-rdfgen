package schema

type HowToStepTrait struct {

}

type HowToStep struct {
	MetaTrait
	HowToStepTrait
	CreativeWorkTrait
	ThingTrait
	ListItemTrait
	IntangibleTrait
	ItemListTrait
	AdditionalTrait
}

func NewHowToStep() (x HowToStep) {
	x.Type = "http://schema.org/HowToStep"

	return
}
