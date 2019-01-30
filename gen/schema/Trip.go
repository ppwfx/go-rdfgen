package schema

type TripTrait struct {

	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense).
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Trip
	//
	HasPart interface{} `json:"hasPart,omitempty" jsonld:"http://schema.org/hasPart"`

	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Trip
	//
	IsPartOf interface{} `json:"isPartOf,omitempty" jsonld:"http://schema.org/isPartOf"`

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Provider interface{} `json:"provider,omitempty" jsonld:"http://schema.org/provider"`

	// The expected arrival time.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ArrivalTime *DateTime `json:"arrivalTime,omitempty" jsonld:"http://schema.org/arrivalTime"`

	// The expected departure time.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DepartureTime *DateTime `json:"departureTime,omitempty" jsonld:"http://schema.org/departureTime"`

	// Destination(s) ( <a class=\"localLink\" href=\"http://schema.org/Place\">Place</a> ) that make up a trip. For a trip where destination order is important use <a class=\"localLink\" href=\"http://schema.org/ItemList\">ItemList</a> to specify that order (see examples).
	//
	// RangeIncludes:
	// - http://schema.org/Place
	// - http://schema.org/ItemList
	//
	Itinerary interface{} `json:"itinerary,omitempty" jsonld:"http://schema.org/itinerary"`

}

type Trip struct {
	MetaTrait
	TripTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTrip() (x Trip) {
	x.Type = "http://schema.org/Trip"

	return
}
