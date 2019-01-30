package schema

type BroadcastEventTrait struct {

	// The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VideoFormat string `json:"videoFormat,omitempty" jsonld:"http://schema.org/videoFormat"`

	// The event being broadcast such as a sporting event or awards ceremony.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	BroadcastOfEvent *Event `json:"broadcastOfEvent,omitempty" jsonld:"http://schema.org/broadcastOfEvent"`

	// True is the broadcast is of a live event.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsLiveBroadcast bool `json:"isLiveBroadcast,omitempty" jsonld:"http://schema.org/isLiveBroadcast"`

}

type BroadcastEvent struct {
	MetaTrait
	BroadcastEventTrait
	PublicationEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewBroadcastEvent() (x BroadcastEvent) {
	x.Type = "http://schema.org/BroadcastEvent"

	return
}
