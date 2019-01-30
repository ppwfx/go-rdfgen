package schema

type PlaceTrait struct {

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

	// A flag to signal that the item, event, or place is accessible for free.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty" jsonld:"http://schema.org/isAccessibleForFree"`

	// A review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Review *Review `json:"review,omitempty" jsonld:"http://schema.org/review"`

	// Review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Reviews *Review `json:"reviews,omitempty" jsonld:"http://schema.org/reviews"`

	// The total number of individuals that may attend an event or venue.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	MaximumAttendeeCapacity *Integer `json:"maximumAttendeeCapacity,omitempty" jsonld:"http://schema.org/maximumAttendeeCapacity"`

	// The special opening hours of a certain place.<br/><br/>\n\nUse this to explicitly override general opening hours brought in scope by <a class=\"localLink\" href=\"http://schema.org/openingHoursSpecification\">openingHoursSpecification</a> or <a class=\"localLink\" href=\"http://schema.org/openingHours\">openingHours</a>.
	//
	// RangeIncludes:
	// - http://schema.org/OpeningHoursSpecification
	//
	SpecialOpeningHoursSpecification *OpeningHoursSpecification `json:"specialOpeningHoursSpecification,omitempty" jsonld:"http://schema.org/specialOpeningHoursSpecification"`

	// The opening hours of a certain place.
	//
	// RangeIncludes:
	// - http://schema.org/OpeningHoursSpecification
	//
	OpeningHoursSpecification *OpeningHoursSpecification `json:"openingHoursSpecification,omitempty" jsonld:"http://schema.org/openingHoursSpecification"`

	// Physical address of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PostalAddress
	//
	Address interface{} `json:"address,omitempty" jsonld:"http://schema.org/address"`

	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.<br/><br/>\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// RangeIncludes:
	// - http://schema.org/PropertyValue
	//
	AdditionalProperty *PropertyValue `json:"additionalProperty,omitempty" jsonld:"http://schema.org/additionalProperty"`

	// The fax number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FaxNumber string `json:"faxNumber,omitempty" jsonld:"http://schema.org/faxNumber"`

	// The telephone number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Telephone string `json:"telephone,omitempty" jsonld:"http://schema.org/telephone"`

	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	//
	// RangeIncludes:
	// - http://schema.org/LocationFeatureSpecification
	//
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty" jsonld:"http://schema.org/amenityFeature"`

	// A short textual code (also called \"store code\") that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.<br/><br/>\n\nFor example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code \"3047\" is a branchCode for a particular branch.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BranchCode string `json:"branchCode,omitempty" jsonld:"http://schema.org/branchCode"`

	// The basic containment relation between a place and one that contains it.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ContainedIn *Place `json:"containedIn,omitempty" jsonld:"http://schema.org/containedIn"`

	// The basic containment relation between a place and one that contains it.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ContainedInPlace *Place `json:"containedInPlace,omitempty" jsonld:"http://schema.org/containedInPlace"`

	// The basic containment relation between a place and another that it contains.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ContainsPlace *Place `json:"containsPlace,omitempty" jsonld:"http://schema.org/containsPlace"`

	// Upcoming or past event associated with this place, organization, or action.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Event *Event `json:"event,omitempty" jsonld:"http://schema.org/event"`

	// Upcoming or past events associated with this place or organization.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	Events *Event `json:"events,omitempty" jsonld:"http://schema.org/events"`

	// The geo coordinates of the place.
	//
	// RangeIncludes:
	// - http://schema.org/GeoCoordinates
	// - http://schema.org/GeoShape
	//
	Geo interface{} `json:"geo,omitempty" jsonld:"http://schema.org/geo"`

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

	// The <a href=\"http://www.gs1.org/gln\">Global Location Number</a> (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	GlobalLocationNumber string `json:"globalLocationNumber,omitempty" jsonld:"http://schema.org/globalLocationNumber"`

	// A URL to a map of the place.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/Map
	//
	HasMap interface{} `json:"hasMap,omitempty" jsonld:"http://schema.org/hasMap"`

	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IsicV4 string `json:"isicV4,omitempty" jsonld:"http://schema.org/isicV4"`

	// An associated logo.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/URL
	//
	Logo interface{} `json:"logo,omitempty" jsonld:"http://schema.org/logo"`

	// A URL to a map of the place.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	Map string `json:"map,omitempty" jsonld:"http://schema.org/map"`

	// A URL to a map of the place.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	Maps string `json:"maps,omitempty" jsonld:"http://schema.org/maps"`

	// A photograph of this place.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/Photograph
	//
	Photo interface{} `json:"photo,omitempty" jsonld:"http://schema.org/photo"`

	// Photographs of this place.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/Photograph
	//
	Photos interface{} `json:"photos,omitempty" jsonld:"http://schema.org/photos"`

	// A flag to signal that the <a class=\"localLink\" href=\"http://schema.org/Place\">Place</a> is open to public visitors.  If this property is omitted there is no assumed default boolean value
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	PublicAccess bool `json:"publicAccess,omitempty" jsonld:"http://schema.org/publicAccess"`

	// Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	SmokingAllowed bool `json:"smokingAllowed,omitempty" jsonld:"http://schema.org/smokingAllowed"`

}

type Place struct {
	MetaTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPlace() (x Place) {
	x.Type = "http://schema.org/Place"

	return
}
