package schema

type InsertActionTrait struct {

	// A sub property of location. The final location of the object or the agent after the action.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ToLocation *Place `json:"toLocation,omitempty" jsonld:"http://schema.org/toLocation"`

}

type InsertAction struct {
	MetaTrait
	InsertActionTrait
	AddActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewInsertAction() (x InsertAction) {
	x.Type = "http://schema.org/InsertAction"

	return
}
