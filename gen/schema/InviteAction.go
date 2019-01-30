package schema

type InviteActionTrait struct {

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

}

type InviteAction struct {
	MetaTrait
	InviteActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewInviteAction() (x InviteAction) {
	x.Type = "http://schema.org/InviteAction"

	return
}
