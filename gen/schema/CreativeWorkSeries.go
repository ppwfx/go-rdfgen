package schema

type CreativeWorkSeriesTrait struct {

	// The end date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	EndDate interface{} `json:"endDate,omitempty" jsonld:"http://schema.org/endDate"`

	// The start date and time of the item (in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>).
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	StartDate interface{} `json:"startDate,omitempty" jsonld:"http://schema.org/startDate"`

	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Issn string `json:"issn,omitempty" jsonld:"http://schema.org/issn"`

}

type CreativeWorkSeries struct {
	MetaTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewCreativeWorkSeries() (x CreativeWorkSeries) {
	x.Type = "http://schema.org/CreativeWorkSeries"

	return
}
