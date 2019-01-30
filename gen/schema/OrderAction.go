package schema

type OrderActionTrait struct {

	// A sub property of instrument. The method of delivery.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty" jsonld:"http://schema.org/deliveryMethod"`

}

type OrderAction struct {
	MetaTrait
	OrderActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewOrderAction() (x OrderAction) {
	x.Type = "http://schema.org/OrderAction"

	return
}
