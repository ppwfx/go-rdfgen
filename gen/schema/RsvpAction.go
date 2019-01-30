package schema

type RsvpActionTrait struct {

	// Comments, typically from users.
	//
	// RangeIncludes:
	// - http://schema.org/Comment
	//
	Comment *Comment `json:"comment,omitempty" jsonld:"http://schema.org/comment"`

	// If responding yes, the number of guests who will attend in addition to the invitee.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	AdditionalNumberOfGuests float64 `json:"additionalNumberOfGuests,omitempty" jsonld:"http://schema.org/additionalNumberOfGuests"`

	// The response (yes, no, maybe) to the RSVP.
	//
	// RangeIncludes:
	// - http://schema.org/RsvpResponseType
	//
	RsvpResponse *RsvpResponseType `json:"rsvpResponse,omitempty" jsonld:"http://schema.org/rsvpResponse"`

}

type RsvpAction struct {
	MetaTrait
	RsvpActionTrait
	InformActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewRsvpAction() (x RsvpAction) {
	x.Type = "http://schema.org/RsvpAction"

	return
}
