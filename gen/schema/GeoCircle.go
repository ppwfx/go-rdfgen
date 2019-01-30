package schema

type GeoCircleTrait struct {

	// Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Distance
	// - http://schema.org/Number
	//
	GeoRadius interface{} `json:"geoRadius,omitempty" jsonld:"http://schema.org/geoRadius"`

	// Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	//
	// RangeIncludes:
	// - http://schema.org/GeoCoordinates
	//
	GeoMidpoint *GeoCoordinates `json:"geoMidpoint,omitempty" jsonld:"http://schema.org/geoMidpoint"`

}

type GeoCircle struct {
	MetaTrait
	GeoCircleTrait
	GeoShapeTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGeoCircle() (x GeoCircle) {
	x.Type = "http://schema.org/GeoCircle"

	return
}
