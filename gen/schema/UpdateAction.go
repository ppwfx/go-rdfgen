package schema

type UpdateActionTrait struct {

	// A sub property of object. The collection target of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	TargetCollection *Thing `json:"targetCollection,omitempty" jsonld:"http://schema.org/targetCollection"`

	// A sub property of object. The collection target of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Collection *Thing `json:"collection,omitempty" jsonld:"http://schema.org/collection"`

}

type UpdateAction struct {
	MetaTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewUpdateAction() (x UpdateAction) {
	x.Type = "http://schema.org/UpdateAction"

	return
}
