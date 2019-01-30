package schema

type DayOfWeekTrait struct {

}

type DayOfWeek struct {
	MetaTrait
	DayOfWeekTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDayOfWeek() (x DayOfWeek) {
	x.Type = "http://schema.org/DayOfWeek"

	return
}
