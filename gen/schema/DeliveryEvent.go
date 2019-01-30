package schema

type DeliveryEventTrait struct {

	// Method used for delivery or shipping.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	HasDeliveryMethod *DeliveryMethod `json:"hasDeliveryMethod,omitempty" jsonld:"http://schema.org/hasDeliveryMethod"`

	// When the item is available for pickup from the store, locker, etc.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailableFrom *DateTime `json:"availableFrom,omitempty" jsonld:"http://schema.org/availableFrom"`

	// After this date, the item will no longer be available for pickup.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailableThrough *DateTime `json:"availableThrough,omitempty" jsonld:"http://schema.org/availableThrough"`

	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccessCode string `json:"accessCode,omitempty" jsonld:"http://schema.org/accessCode"`

}

type DeliveryEvent struct {
	MetaTrait
	DeliveryEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewDeliveryEvent() (x DeliveryEvent) {
	x.Type = "http://schema.org/DeliveryEvent"

	return
}
