package schema

type EventTrait struct {

	// The subject matter of the content.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	About *Thing `json:"about,omitempty" jsonld:"http://schema.org/about"`

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

	// An intended audience, i.e. a group for whom something was created.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	Audience *Audience `json:"audience,omitempty" jsonld:"http://schema.org/audience"`

	// A secondary contributor to the CreativeWork or Event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Contributor interface{} `json:"contributor,omitempty" jsonld:"http://schema.org/contributor"`

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Funder interface{} `json:"funder,omitempty" jsonld:"http://schema.org/funder"`

	// The language of the content or performance or used in an action. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/availableLanguage\">availableLanguage</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Language
	// - http://schema.org/Text
	//
	InLanguage interface{} `json:"inLanguage,omitempty" jsonld:"http://schema.org/inLanguage"`

	// A flag to signal that the item, event, or place is accessible for free.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty" jsonld:"http://schema.org/isAccessibleForFree"`

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// A review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Review *Review `json:"review,omitempty" jsonld:"http://schema.org/review"`

	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Sponsor interface{} `json:"sponsor,omitempty" jsonld:"http://schema.org/sponsor"`

	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Translator interface{} `json:"translator,omitempty" jsonld:"http://schema.org/translator"`

	// The typical expected age range, e.g. '7-9', '11-'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TypicalAgeRange string `json:"typicalAgeRange,omitempty" jsonld:"http://schema.org/typicalAgeRange"`

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Actor *Person `json:"actor,omitempty" jsonld:"http://schema.org/actor"`

	// A person or organization attending the event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Attendee interface{} `json:"attendee,omitempty" jsonld:"http://schema.org/attendee"`

	// A person attending the event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Attendees interface{} `json:"attendees,omitempty" jsonld:"http://schema.org/attendees"`

	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Composer interface{} `json:"composer,omitempty" jsonld:"http://schema.org/composer"`

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Director *Person `json:"director,omitempty" jsonld:"http://schema.org/director"`

	// The time admission will commence.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DoorTime *DateTime `json:"doorTime,omitempty" jsonld:"http://schema.org/doorTime"`

	// The end date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	EndDate interface{} `json:"endDate,omitempty" jsonld:"http://schema.org/endDate"`

	// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	//
	// RangeIncludes:
	// - http://schema.org/EventStatusType
	//
	EventStatus *EventStatusType `json:"eventStatus,omitempty" jsonld:"http://schema.org/eventStatus"`

	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	// - http://schema.org/PostalAddress
	// - http://schema.org/Text
	//
	Location interface{} `json:"location,omitempty" jsonld:"http://schema.org/location"`

	// The total number of individuals that may attend an event or venue.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	MaximumAttendeeCapacity *Integer `json:"maximumAttendeeCapacity,omitempty" jsonld:"http://schema.org/maximumAttendeeCapacity"`

	// An organizer of an Event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Organizer interface{} `json:"organizer,omitempty" jsonld:"http://schema.org/organizer"`

	// A performer at the event&#x2014;for example, a presenter, musician, musical group or actor.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Performer interface{} `json:"performer,omitempty" jsonld:"http://schema.org/performer"`

	// The main performer or performers of the event&#x2014;for example, a presenter, musician, or actor.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Performers interface{} `json:"performers,omitempty" jsonld:"http://schema.org/performers"`

	// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	PreviousStartDate *Date `json:"previousStartDate,omitempty" jsonld:"http://schema.org/previousStartDate"`

	// The CreativeWork that captured all or part of this Event.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	RecordedIn *CreativeWork `json:"recordedIn,omitempty" jsonld:"http://schema.org/recordedIn"`

	// The number of attendee places for an event that remain unallocated.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	RemainingAttendeeCapacity *Integer `json:"remainingAttendeeCapacity,omitempty" jsonld:"http://schema.org/remainingAttendeeCapacity"`

	// The start date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	StartDate interface{} `json:"startDate,omitempty" jsonld:"http://schema.org/startDate"`

	// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	SubEvent *Event `json:"subEvent,omitempty" jsonld:"http://schema.org/subEvent"`

	// Events that are a part of this event. For example, a conference event includes many presentations, each subEvents of the conference.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	SubEvents *Event `json:"subEvents,omitempty" jsonld:"http://schema.org/subEvents"`

	// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	SuperEvent *Event `json:"superEvent,omitempty" jsonld:"http://schema.org/superEvent"`

	// A work featured in some event, e.g. exhibited in an ExhibitionEvent.\n       Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	WorkFeatured *CreativeWork `json:"workFeatured,omitempty" jsonld:"http://schema.org/workFeatured"`

	// A work performed in some event, for example a play performed in a TheaterEvent.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	WorkPerformed *CreativeWork `json:"workPerformed,omitempty" jsonld:"http://schema.org/workPerformed"`

}

type Event struct {
	MetaTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewEvent() (x Event) {
	x.Type = "http://schema.org/Event"

	return
}
