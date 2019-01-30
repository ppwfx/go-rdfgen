package schema

type AggregateOfferTrait struct {

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// The highest price of all offers available.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	HighPrice interface{} `json:"highPrice,omitempty" jsonld:"http://schema.org/highPrice"`

	// The lowest price of all offers available.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	LowPrice interface{} `json:"lowPrice,omitempty" jsonld:"http://schema.org/lowPrice"`

	// The number of offers for the product.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	OfferCount *Integer `json:"offerCount,omitempty" jsonld:"http://schema.org/offerCount"`

}

type AggregateOffer struct {
	MetaTrait
	AggregateOfferTrait
	OfferTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewAggregateOffer() (x AggregateOffer) {
	x.Type = "http://schema.org/AggregateOffer"

	return
}
