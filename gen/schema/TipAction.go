package schema

type TipActionTrait struct {

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

type TipAction struct {
	MetaTrait
	TipActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTipAction() (x TipAction) {
	x.Type = "http://schema.org/TipAction"

	return
}
