package schema

type PublicationEventTrait struct {

	// A flag to signal that the item, event, or place is accessible for free.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty" jsonld:"http://schema.org/isAccessibleForFree"`

	// A flag to signal that the item, event, or place is accessible for free.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	Free bool `json:"free,omitempty" jsonld:"http://schema.org/free"`

	// An agent associated with the publication event.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	PublishedBy interface{} `json:"publishedBy,omitempty" jsonld:"http://schema.org/publishedBy"`

	// A broadcast service associated with the publication event.
	//
	// RangeIncludes:
	// - http://schema.org/BroadcastService
	//
	PublishedOn *BroadcastService `json:"publishedOn,omitempty" jsonld:"http://schema.org/publishedOn"`

}

type PublicationEvent struct {
	MetaTrait
	PublicationEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewPublicationEvent() (x PublicationEvent) {
	x.Type = "http://schema.org/PublicationEvent"

	return
}
