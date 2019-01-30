package schema

type AuthorizeActionTrait struct {

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

type AuthorizeAction struct {
	MetaTrait
	AuthorizeActionTrait
	AllocateActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAuthorizeAction() (x AuthorizeAction) {
	x.Type = "http://schema.org/AuthorizeAction"

	return
}
