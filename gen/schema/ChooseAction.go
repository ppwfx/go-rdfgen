package schema

type ChooseActionTrait struct {

	// A sub property of object. The options subject to this action.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	//
	ActionOption interface{} `json:"actionOption,omitempty" jsonld:"http://schema.org/actionOption"`

	// A sub property of object. The options subject to this action.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	//
	Option interface{} `json:"option,omitempty" jsonld:"http://schema.org/option"`

}

type ChooseAction struct {
	MetaTrait
	ChooseActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewChooseAction() (x ChooseAction) {
	x.Type = "http://schema.org/ChooseAction"

	return
}
