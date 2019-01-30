package schema

type InformActionTrait struct {

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

}

type InformAction struct {
	MetaTrait
	InformActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewInformAction() (x InformAction) {
	x.Type = "http://schema.org/InformAction"

	return
}
