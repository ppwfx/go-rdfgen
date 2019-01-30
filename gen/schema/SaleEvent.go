package schema

type SaleEventTrait struct {

}

type SaleEvent struct {
	MetaTrait
	SaleEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewSaleEvent() (x SaleEvent) {
	x.Type = "http://schema.org/SaleEvent"

	return
}
