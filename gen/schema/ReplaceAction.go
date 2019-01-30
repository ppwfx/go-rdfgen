package schema

type ReplaceActionTrait struct {

	// A sub property of object. The object that replaces.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Replacer *Thing `json:"replacer,omitempty" jsonld:"http://schema.org/replacer"`

	// A sub property of object. The object that is being replaced.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Replacee *Thing `json:"replacee,omitempty" jsonld:"http://schema.org/replacee"`

}

type ReplaceAction struct {
	MetaTrait
	ReplaceActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReplaceAction() (x ReplaceAction) {
	x.Type = "http://schema.org/ReplaceAction"

	return
}
