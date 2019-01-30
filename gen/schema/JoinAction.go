package schema

type JoinActionTrait struct {

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

}

type JoinAction struct {
	MetaTrait
	JoinActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewJoinAction() (x JoinAction) {
	x.Type = "http://schema.org/JoinAction"

	return
}
