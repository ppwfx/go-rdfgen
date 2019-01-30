package schema

type GeospatialGeometryTrait struct {

	// Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. \"a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a\". As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyContains interface{} `json:"geospatiallyContains,omitempty" jsonld:"http://schema.org/geospatiallyContains"`

	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyCoveredBy interface{} `json:"geospatiallyCoveredBy,omitempty" jsonld:"http://schema.org/geospatiallyCoveredBy"`

	// Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. \"Every point of b is a point of (the interior or boundary of) a\". As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyCovers interface{} `json:"geospatiallyCovers,omitempty" jsonld:"http://schema.org/geospatiallyCovers"`

	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: \"a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them\". As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyCrosses interface{} `json:"geospatiallyCrosses,omitempty" jsonld:"http://schema.org/geospatiallyCrosses"`

	// Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.\" (a symmetric relationship, as defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>)
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyDisjoint interface{} `json:"geospatiallyDisjoint,omitempty" jsonld:"http://schema.org/geospatiallyDisjoint"`

	// Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>. \"Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other\" (a symmetric relationship)
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyEquals interface{} `json:"geospatiallyEquals,omitempty" jsonld:"http://schema.org/geospatiallyEquals"`

	// Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyIntersects interface{} `json:"geospatiallyIntersects,omitempty" jsonld:"http://schema.org/geospatiallyIntersects"`

	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyOverlaps interface{} `json:"geospatiallyOverlaps,omitempty" jsonld:"http://schema.org/geospatiallyOverlaps"`

	// Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.\" (a symmetric relationship, as defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a> )
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyTouches interface{} `json:"geospatiallyTouches,omitempty" jsonld:"http://schema.org/geospatiallyTouches"`

	// Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in <a href=\"https://en.wikipedia.org/wiki/DE-9IM\">DE-9IM</a>.
	//
	// RangeIncludes:
	// - http://schema.org/GeospatialGeometry
	// - http://schema.org/Place
	//
	GeospatiallyWithin interface{} `json:"geospatiallyWithin,omitempty" jsonld:"http://schema.org/geospatiallyWithin"`

}

type GeospatialGeometry struct {
	MetaTrait
	GeospatialGeometryTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGeospatialGeometry() (x GeospatialGeometry) {
	x.Type = "http://schema.org/GeospatialGeometry"

	return
}
