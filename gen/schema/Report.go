package schema

type ReportTrait struct {

	// The number or other unique designator assigned to a Report by the publishing organization.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ReportNumber string `json:"reportNumber,omitempty" jsonld:"http://schema.org/reportNumber"`

}

type Report struct {
	MetaTrait
	ReportTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewReport() (x Report) {
	x.Type = "http://schema.org/Report"

	return
}
