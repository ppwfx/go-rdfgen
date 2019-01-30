package schema

type PlayActionTrait struct {

	// An intended audience, i.e. a group for whom something was created.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	Audience *Audience `json:"audience,omitempty" jsonld:"http://schema.org/audience"`

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

}

type PlayAction struct {
	MetaTrait
	PlayActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPlayAction() (x PlayAction) {
	x.Type = "http://schema.org/PlayAction"

	return
}
