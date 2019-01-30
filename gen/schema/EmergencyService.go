package schema

type EmergencyServiceTrait struct {

}

type EmergencyService struct {
	MetaTrait
	EmergencyServiceTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewEmergencyService() (x EmergencyService) {
	x.Type = "http://schema.org/EmergencyService"

	return
}
