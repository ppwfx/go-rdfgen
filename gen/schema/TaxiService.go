package schema

type TaxiServiceTrait struct {

}

type TaxiService struct {
	MetaTrait
	TaxiServiceTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTaxiService() (x TaxiService) {
	x.Type = "http://schema.org/TaxiService"

	return
}
