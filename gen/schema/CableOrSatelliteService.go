package schema

type CableOrSatelliteServiceTrait struct {

}

type CableOrSatelliteService struct {
	MetaTrait
	CableOrSatelliteServiceTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewCableOrSatelliteService() (x CableOrSatelliteService) {
	x.Type = "http://schema.org/CableOrSatelliteService"

	return
}
