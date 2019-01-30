package schema

type LeaveActionTrait struct {

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

}

type LeaveAction struct {
	MetaTrait
	LeaveActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewLeaveAction() (x LeaveAction) {
	x.Type = "http://schema.org/LeaveAction"

	return
}
