package schema

type TaxiTrait struct {

}

type Taxi struct {
	MetaTrait
	TaxiTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTaxi() (x Taxi) {
	x.Type = "http://schema.org/Taxi"

	return
}
