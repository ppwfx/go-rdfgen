package schema

type TransferActionTrait struct {

	// A sub property of location. The original location of the object or the agent before the action.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	FromLocation *Place `json:"fromLocation,omitempty" jsonld:"http://schema.org/fromLocation"`

	// A sub property of location. The final location of the object or the agent after the action.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ToLocation *Place `json:"toLocation,omitempty" jsonld:"http://schema.org/toLocation"`

}

type TransferAction struct {
	MetaTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTransferAction() (x TransferAction) {
	x.Type = "http://schema.org/TransferAction"

	return
}
