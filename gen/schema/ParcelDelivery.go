package schema

type ParcelDeliveryTrait struct {

	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Provider interface{} `json:"provider,omitempty" jsonld:"http://schema.org/provider"`

	// 'carrier' is an out-dated term indicating the 'provider' for parcel delivery and flights.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Carrier *Organization `json:"carrier,omitempty" jsonld:"http://schema.org/carrier"`

	// Destination address.
	//
	// RangeIncludes:
	// - http://schema.org/PostalAddress
	//
	DeliveryAddress *PostalAddress `json:"deliveryAddress,omitempty" jsonld:"http://schema.org/deliveryAddress"`

	// New entry added as the package passes through each leg of its journey (from shipment to final delivery).
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryEvent
	//
	DeliveryStatus *DeliveryEvent `json:"deliveryStatus,omitempty" jsonld:"http://schema.org/deliveryStatus"`

	// The earliest date the package may arrive.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ExpectedArrivalFrom *DateTime `json:"expectedArrivalFrom,omitempty" jsonld:"http://schema.org/expectedArrivalFrom"`

	// The latest date the package may arrive.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ExpectedArrivalUntil *DateTime `json:"expectedArrivalUntil,omitempty" jsonld:"http://schema.org/expectedArrivalUntil"`

	// Method used for delivery or shipping.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty" jsonld:"http://schema.org/hasDeliveryMethod"`

	// Item(s) being shipped.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	//
	ItemShipped *Product `json:"itemShipped,omitempty" jsonld:"http://schema.org/itemShipped"`

	// Shipper's address.
	//
	// RangeIncludes:
	// - http://schema.org/PostalAddress
	//
	OriginAddress *PostalAddress `json:"originAddress,omitempty" jsonld:"http://schema.org/originAddress"`

	// The overall order the items in this delivery were included in.
	//
	// RangeIncludes:
	// - http://schema.org/Order
	//
	PartOfOrder *Order `json:"partOfOrder,omitempty" jsonld:"http://schema.org/partOfOrder"`

	// Shipper tracking number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TrackingNumber string `json:"trackingNumber,omitempty" jsonld:"http://schema.org/trackingNumber"`

	// Tracking url for the parcel delivery.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	TrackingUrl string `json:"trackingUrl,omitempty" jsonld:"http://schema.org/trackingUrl"`

}

type ParcelDelivery struct {
	MetaTrait
	ParcelDeliveryTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewParcelDelivery() (x ParcelDelivery) {
	x.Type = "http://schema.org/ParcelDelivery"

	return
}
