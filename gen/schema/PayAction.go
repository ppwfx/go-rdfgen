package schema

type PayActionTrait struct {

	// A sub property of participant. The participant who is at the receiving end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Audience
	// - http://schema.org/Organization
	//
	Recipient interface{} `json:"recipient,omitempty" jsonld:"http://schema.org/recipient"`

	// A goal towards an action is taken. Can be concrete or abstract.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	// - http://schema.org/MedicalDevicePurpose
	//
	Purpose interface{} `json:"purpose,omitempty" jsonld:"http://schema.org/purpose"`

}

type PayAction struct {
	MetaTrait
	PayActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPayAction() (x PayAction) {
	x.Type = "http://schema.org/PayAction"

	return
}
