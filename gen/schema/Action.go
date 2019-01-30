package schema

type ActionTrait struct {

	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	// - http://schema.org/PostalAddress
	// - http://schema.org/Text
	//
	Location interface{} `json:"location,omitempty" jsonld:"http://schema.org/location"`

	// Indicates the current disposition of the Action.
	//
	// RangeIncludes:
	// - http://schema.org/ActionStatusType
	//
	ActionStatus *ActionStatusType `json:"actionStatus,omitempty" jsonld:"http://schema.org/actionStatus"`

	// The direct performer or driver of the action (animate or inanimate). e.g. <em>John</em> wrote a book.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Agent interface{} `json:"agent,omitempty" jsonld:"http://schema.org/agent"`

	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to <em>December</em>.<br/><br/>\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	EndTime *DateTime `json:"endTime,omitempty" jsonld:"http://schema.org/endTime"`

	// For failed actions, more information on the cause of the failure.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Error *Thing `json:"error,omitempty" jsonld:"http://schema.org/error"`

	// The object that helped the agent perform the action. e.g. John wrote a book with <em>a pen</em>.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Instrument *Thing `json:"instrument,omitempty" jsonld:"http://schema.org/instrument"`

	// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn't). e.g. John read <em>a book</em>.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Object *Thing `json:"object,omitempty" jsonld:"http://schema.org/object"`

	// Other co-agents that participated in the action indirectly. e.g. John wrote a book with <em>Steve</em>.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Participant interface{} `json:"participant,omitempty" jsonld:"http://schema.org/participant"`

	// The result produced in the action. e.g. John wrote <em>a book</em>.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Result *Thing `json:"result,omitempty" jsonld:"http://schema.org/result"`

	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from <em>January</em> to December.<br/><br/>\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	StartTime *DateTime `json:"startTime,omitempty" jsonld:"http://schema.org/startTime"`

	// Indicates a target EntryPoint for an Action.
	//
	// RangeIncludes:
	// - http://schema.org/EntryPoint
	//
	Target *EntryPoint `json:"target,omitempty" jsonld:"http://schema.org/target"`

}

type Action struct {
	MetaTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAction() (x Action) {
	x.Type = "http://schema.org/Action"

	return
}
