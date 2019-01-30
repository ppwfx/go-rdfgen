package schema

type ConsumeActionTrait struct {

	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty" jsonld:"http://schema.org/expectsAcceptanceOf"`

	// A set of requirements that a must be fulfilled in order to perform an Action. If more than one value is specied, fulfilling one set of requirements will allow the Action to be performed.
	//
	// RangeIncludes:
	// - http://schema.org/ActionAccessSpecification
	//
	ActionAccessibilityRequirement *ActionAccessSpecification `json:"actionAccessibilityRequirement,omitempty" jsonld:"http://schema.org/actionAccessibilityRequirement"`

}

type ConsumeAction struct {
	MetaTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewConsumeAction() (x ConsumeAction) {
	x.Type = "http://schema.org/ConsumeAction"

	return
}
