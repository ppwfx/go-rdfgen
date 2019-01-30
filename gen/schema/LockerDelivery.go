package schema

type LockerDeliveryTrait struct {

}

type LockerDelivery struct {
	MetaTrait
	LockerDeliveryTrait
	DeliveryMethodTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLockerDelivery() (x LockerDelivery) {
	x.Type = "http://schema.org/LockerDelivery"

	return
}
