package schema

type ScheduleTrait struct {

	// Defines the day(s) of the week on which a recurring <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> takes place
	//
	// RangeIncludes:
	// - http://schema.org/DayOfWeek
	//
	ByDay *DayOfWeek `json:"byDay,omitempty" jsonld:"http://schema.org/byDay"`

	// Defines the month(s) of the year on which a recurring <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> takes place. Specified as an <a class=\"localLink\" href=\"http://schema.org/Integer\">Integer</a> between 1-12. January is 1.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	ByMonth *Integer `json:"byMonth,omitempty" jsonld:"http://schema.org/byMonth"`

	// Defines the day(s) of the month on which a recurring <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> takes place. Specified as an <a class=\"localLink\" href=\"http://schema.org/Integer\">Integer</a> between 1-31.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	ByMonthDay *Integer `json:"byMonthDay,omitempty" jsonld:"http://schema.org/byMonthDay"`

	// Associates an <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> with a <a class=\"localLink\" href=\"http://schema.org/Schedule\">Schedule</a>. There are circumstances where it is preferable to share a schedule for a series of\n      repeating events rather than data on the individual events themselves. For example, a website or application might prefer to publish a schedule for a weekly\n      gym class rather than provide data on every event. A schedule could be processed by applications to add forthcoming events to a calendar. An <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> that\n      is associated with a <a class=\"localLink\" href=\"http://schema.org/Schedule\">Schedule</a> using this property should not have <a class=\"localLink\" href=\"http://schema.org/startDate\">startDate</a> or <a class=\"localLink\" href=\"http://schema.org/endDate\">endDate</a> properties. These are instead defined within the associated\n      <a class=\"localLink\" href=\"http://schema.org/Schedule\">Schedule</a>, this avoids any ambiguity for clients using the data. The propery might have repeated values to specify different schedules, e.g. for different months\n      or seasons.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	EventSchedule *Duration `json:"eventSchedule,omitempty" jsonld:"http://schema.org/eventSchedule"`

	// Defines a <a class=\"localLink\" href=\"http://schema.org/Date\">Date</a> or <a class=\"localLink\" href=\"http://schema.org/DateTime\">DateTime</a> during which a scheduled <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> will not take place. The property allows exceptions to\n      a <a class=\"localLink\" href=\"http://schema.org/Schedule\">Schedule</a> to be specified. If an exception is specified as a <a class=\"localLink\" href=\"http://schema.org/DateTime\">DateTime</a> then only the event that would have started at that specific date and time\n      should be excluded from the schedule. If an exception is specified as a <a class=\"localLink\" href=\"http://schema.org/Date\">Date</a> then any event that is scheduled for that 24 hour period should be\n      excluded from the schedule. This allows a whole day to be excluded from the schedule without having to itemise every scheduled event.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	ExceptDate interface{} `json:"exceptDate,omitempty" jsonld:"http://schema.org/exceptDate"`

	// Defines the number of times a recurring <a class=\"localLink\" href=\"http://schema.org/Event\">Event</a> will take place
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	RepeatCount *Integer `json:"repeatCount,omitempty" jsonld:"http://schema.org/repeatCount"`

	// Defines the frequency at which <a class=\"localLink\" href=\"http://schema.org/Events\">Events</a> will occur according to a schedule <a class=\"localLink\" href=\"http://schema.org/Schedule\">Schedule</a>. The intervals between\n      events should be defined as a <a class=\"localLink\" href=\"http://schema.org/Duration\">Duration</a> of time.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Duration
	//
	RepeatFrequency interface{} `json:"repeatFrequency,omitempty" jsonld:"http://schema.org/repeatFrequency"`

}

type Schedule struct {
	MetaTrait
	ScheduleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSchedule() (x Schedule) {
	x.Type = "http://schema.org/Schedule"

	return
}
