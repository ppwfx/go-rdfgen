package schema

type GiveActionTrait struct {

	// A sub property of participant. The participant who is at the receiving end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Audience
	// - http://schema.org/Organization
	//
	Recipient interface{} `json:"recipient,omitempty" jsonld:"http://schema.org/recipient"`

}

type GiveAction struct {
	MetaTrait
	GiveActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewGiveAction() (x GiveAction) {
	x.Type = "http://schema.org/GiveAction"

	return
}
