package schema

type OrderStatusTrait struct {

}

type OrderStatus struct {
	MetaTrait
	OrderStatusTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOrderStatus() (x OrderStatus) {
	x.Type = "http://schema.org/OrderStatus"

	return
}
