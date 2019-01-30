package schema

type HowToSectionTrait struct {

	// A single step item (as HowToStep, text, document, video, etc.) or a HowToSection (originally misnamed 'steps'; 'step' is preferred).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/ItemList
	//
	Steps interface{} `json:"steps,omitempty" jsonld:"http://schema.org/steps"`

}

type HowToSection struct {
	MetaTrait
	HowToSectionTrait
	CreativeWorkTrait
	ThingTrait
	ListItemTrait
	IntangibleTrait
	ItemListTrait
	AdditionalTrait
}

func NewHowToSection() (x HowToSection) {
	x.Type = "http://schema.org/HowToSection"

	return
}
