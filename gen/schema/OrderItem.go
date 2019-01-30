package schema

type OrderItemTrait struct {

	// The delivery of the parcel related to this order or order item.
	//
	// RangeIncludes:
	// - http://schema.org/ParcelDelivery
	//
	OrderDelivery *ParcelDelivery `json:"orderDelivery,omitempty" jsonld:"http://schema.org/orderDelivery"`

	// The item ordered.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/OrderItem
	//
	OrderedItem interface{} `json:"orderedItem,omitempty" jsonld:"http://schema.org/orderedItem"`

	// The identifier of the order item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	OrderItemNumber string `json:"orderItemNumber,omitempty" jsonld:"http://schema.org/orderItemNumber"`

	// The current status of the order item.
	//
	// RangeIncludes:
	// - http://schema.org/OrderStatus
	//
	OrderItemStatus *OrderStatus `json:"orderItemStatus,omitempty" jsonld:"http://schema.org/orderItemStatus"`

	// The number of the item ordered. If the property is not set, assume the quantity is one.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	OrderQuantity float64 `json:"orderQuantity,omitempty" jsonld:"http://schema.org/orderQuantity"`

}

type OrderItem struct {
	MetaTrait
	OrderItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOrderItem() (x OrderItem) {
	x.Type = "http://schema.org/OrderItem"

	return
}
