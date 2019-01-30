package schema

type ParcelServiceTrait struct {

}

type ParcelService struct {
	MetaTrait
	ParcelServiceTrait
	DeliveryMethodTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewParcelService() (x ParcelService) {
	x.Type = "http://schema.org/ParcelService"

	return
}
