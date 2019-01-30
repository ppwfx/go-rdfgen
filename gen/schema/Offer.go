package schema

type OfferTrait struct {

	// The date when the item becomes valid.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidFrom *DateTime `json:"validFrom,omitempty" jsonld:"http://schema.org/validFrom"`

	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidThrough *DateTime `json:"validThrough,omitempty" jsonld:"http://schema.org/validThrough"`

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

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

	// The payment method(s) accepted by seller for this offer.
	//
	// RangeIncludes:
	// - http://schema.org/LoanOrCredit
	// - http://schema.org/PaymentMethod
	//
	AcceptedPaymentMethod interface{} `json:"acceptedPaymentMethod,omitempty" jsonld:"http://schema.org/acceptedPaymentMethod"`

	// The amount of time that is required between accepting the offer and the actual usage of the resource or service.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	AdvanceBookingRequirement *QuantitativeValue `json:"advanceBookingRequirement,omitempty" jsonld:"http://schema.org/advanceBookingRequirement"`

	// The geographic area where a service or offered item is provided.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	AreaServed interface{} `json:"areaServed,omitempty" jsonld:"http://schema.org/areaServed"`

	// The availability of this item&#x2014;for example In stock, Out of stock, Pre-order, etc.
	//
	// RangeIncludes:
	// - http://schema.org/ItemAvailability
	//
	Availability *ItemAvailability `json:"availability,omitempty" jsonld:"http://schema.org/availability"`

	// The end of the availability of the product or service included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailabilityEnds *DateTime `json:"availabilityEnds,omitempty" jsonld:"http://schema.org/availabilityEnds"`

	// The beginning of the availability of the product or service included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailabilityStarts *DateTime `json:"availabilityStarts,omitempty" jsonld:"http://schema.org/availabilityStarts"`

	// The place(s) from which the offer can be obtained (e.g. store locations).
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	AvailableAtOrFrom *Place `json:"availableAtOrFrom,omitempty" jsonld:"http://schema.org/availableAtOrFrom"`

	// The delivery method(s) available for this offer.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	AvailableDeliveryMethod *DeliveryMethod `json:"availableDeliveryMethod,omitempty" jsonld:"http://schema.org/availableDeliveryMethod"`

	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	//
	// RangeIncludes:
	// - http://schema.org/BusinessFunction
	//
	BusinessFunction *BusinessFunction `json:"businessFunction,omitempty" jsonld:"http://schema.org/businessFunction"`

	// The typical delay between the receipt of the order and the goods either leaving the warehouse or being prepared for pickup, in case the delivery method is on site pickup.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	DeliveryLeadTime *QuantitativeValue `json:"deliveryLeadTime,omitempty" jsonld:"http://schema.org/deliveryLeadTime"`

	// The type(s) of customers for which the given offer is valid.
	//
	// RangeIncludes:
	// - http://schema.org/BusinessEntityType
	//
	EligibleCustomerType *BusinessEntityType `json:"eligibleCustomerType,omitempty" jsonld:"http://schema.org/eligibleCustomerType"`

	// The duration for which the given offer is valid.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	EligibleDuration *QuantitativeValue `json:"eligibleDuration,omitempty" jsonld:"http://schema.org/eligibleDuration"`

	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	EligibleQuantity *QuantitativeValue `json:"eligibleQuantity,omitempty" jsonld:"http://schema.org/eligibleQuantity"`

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.<br/><br/>\n\nSee also <a class=\"localLink\" href=\"http://schema.org/ineligibleRegion\">ineligibleRegion</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	EligibleRegion interface{} `json:"eligibleRegion,omitempty" jsonld:"http://schema.org/eligibleRegion"`

	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	//
	// RangeIncludes:
	// - http://schema.org/PriceSpecification
	//
	EligibleTransactionVolume *PriceSpecification `json:"eligibleTransactionVolume,omitempty" jsonld:"http://schema.org/eligibleTransactionVolume"`

	// The GTIN-12 code of the product, or the product to which the offer refers. The GTIN-12 is the 12-digit GS1 Identification Key composed of a U.P.C. Company Prefix, Item Reference, and Check Digit used to identify trade items. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin12 string `json:"gtin12,omitempty" jsonld:"http://schema.org/gtin12"`

	// The GTIN-13 code of the product, or the product to which the offer refers. This is equivalent to 13-digit ISBN codes and EAN UCC-13. Former 12-digit UPC codes can be converted into a GTIN-13 code by simply adding a preceeding zero. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin13 string `json:"gtin13,omitempty" jsonld:"http://schema.org/gtin13"`

	// The GTIN-14 code of the product, or the product to which the offer refers. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin14 string `json:"gtin14,omitempty" jsonld:"http://schema.org/gtin14"`

	// The <a href=\"http://apps.gs1.org/GDD/glossary/Pages/GTIN-8.aspx\">GTIN-8</a> code of the product, or the product to which the offer refers. This code is also known as EAN/UCC-8 or 8-digit EAN. See <a href=\"http://www.gs1.org/barcodes/technical/idkeys/gtin\">GS1 GTIN Summary</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Gtin8 string `json:"gtin8,omitempty" jsonld:"http://schema.org/gtin8"`

	// This links to a node or nodes indicating the exact quantity of the products included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/TypeAndQuantityNode
	//
	IncludesObject *TypeAndQuantityNode `json:"includesObject,omitempty" jsonld:"http://schema.org/includesObject"`

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.<br/><br/>\n\nSee also <a class=\"localLink\" href=\"http://schema.org/eligibleRegion\">eligibleRegion</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	IneligibleRegion interface{} `json:"ineligibleRegion,omitempty" jsonld:"http://schema.org/ineligibleRegion"`

	// The current approximate inventory level for the item or items.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	InventoryLevel *QuantitativeValue `json:"inventoryLevel,omitempty" jsonld:"http://schema.org/inventoryLevel"`

	// A predefined value from OfferItemCondition or a textual description of the condition of the product or service, or the products or services included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/OfferItemCondition
	//
	ItemCondition *OfferItemCondition `json:"itemCondition,omitempty" jsonld:"http://schema.org/itemCondition"`

	// The item being offered.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Service
	//
	ItemOffered interface{} `json:"itemOffered,omitempty" jsonld:"http://schema.org/itemOffered"`

	// The Manufacturer Part Number (MPN) of the product, or the product to which the offer refers.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Mpn string `json:"mpn,omitempty" jsonld:"http://schema.org/mpn"`

	// One or more detailed price specifications, indicating the unit price and delivery or payment charges.
	//
	// RangeIncludes:
	// - http://schema.org/PriceSpecification
	//
	PriceSpecification *PriceSpecification `json:"priceSpecification,omitempty" jsonld:"http://schema.org/priceSpecification"`

	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Seller interface{} `json:"seller,omitempty" jsonld:"http://schema.org/seller"`

	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SerialNumber string `json:"serialNumber,omitempty" jsonld:"http://schema.org/serialNumber"`

	// The Stock Keeping Unit (SKU), i.e. a merchant-specific identifier for a product or service, or the product to which the offer refers.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Sku string `json:"sku,omitempty" jsonld:"http://schema.org/sku"`

	// The warranty promise(s) included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/WarrantyPromise
	//
	Warranty *WarrantyPromise `json:"warranty,omitempty" jsonld:"http://schema.org/warranty"`

	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/PhysicalActivityCategory
	//
	Category interface{} `json:"category,omitempty" jsonld:"http://schema.org/category"`

	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.<br/><br/>\n\nUsage guidelines:<br/><br/>\n\n<ul>\n<li>Use the <a class=\"localLink\" href=\"http://schema.org/priceCurrency\">priceCurrency</a> property (with standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\") instead of including <a href=\"http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign\">ambiguous symbols</a> such as '$' in the value.</li>\n<li>Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.</li>\n<li>Note that both <a href=\"http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute\">RDFa</a> and Microdata syntax allow the use of a \"content=\" attribute for publishing simple machine-readable values alongside more human-friendly formatting.</li>\n<li>Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similiar Unicode symbols.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Price interface{} `json:"price,omitempty" jsonld:"http://schema.org/price"`

	// The currency of the price, or a price component when attached to <a class=\"localLink\" href=\"http://schema.org/PriceSpecification\">PriceSpecification</a> and its subtypes.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceCurrency string `json:"priceCurrency,omitempty" jsonld:"http://schema.org/priceCurrency"`

	// An additional offer that can only be obtained in combination with the first base offer (e.g. supplements and extensions that are available for a surcharge).
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	AddOn *Offer `json:"addOn,omitempty" jsonld:"http://schema.org/addOn"`

	// A pointer to the organization or person making the offer.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	OfferedBy interface{} `json:"offeredBy,omitempty" jsonld:"http://schema.org/offeredBy"`

	// The date after which the price is no longer available.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	PriceValidUntil *Date `json:"priceValidUntil,omitempty" jsonld:"http://schema.org/priceValidUntil"`

}

type Offer struct {
	MetaTrait
	OfferTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOffer() (x Offer) {
	x.Type = "http://schema.org/Offer"

	return
}
