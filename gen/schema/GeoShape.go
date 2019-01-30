package schema

type GeoShapeTrait struct {

	// Physical address of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PostalAddress
	//
	Address interface{} `json:"address,omitempty" jsonld:"http://schema.org/address"`

	// The country. For example, USA. You can also provide the two-letter <a href=\"http://en.wikipedia.org/wiki/ISO_3166-1\">ISO 3166-1 alpha-2 country code</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Country
	//
	AddressCountry interface{} `json:"addressCountry,omitempty" jsonld:"http://schema.org/addressCountry"`

	// The elevation of a location (<a href=\"https://en.wikipedia.org/wiki/World_Geodetic_System\">WGS 84</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Elevation interface{} `json:"elevation,omitempty" jsonld:"http://schema.org/elevation"`

	// The postal code. For example, 94043.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PostalCode string `json:"postalCode,omitempty" jsonld:"http://schema.org/postalCode"`

	// A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Box string `json:"box,omitempty" jsonld:"http://schema.org/box"`

	// A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Circle string `json:"circle,omitempty" jsonld:"http://schema.org/circle"`

	// A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Line string `json:"line,omitempty" jsonld:"http://schema.org/line"`

	// A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Polygon string `json:"polygon,omitempty" jsonld:"http://schema.org/polygon"`

}

type GeoShape struct {
	MetaTrait
	GeoShapeTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGeoShape() (x GeoShape) {
	x.Type = "http://schema.org/GeoShape"

	return
}
