package schema

type AutoRepairTrait struct {

}

type AutoRepair struct {
	MetaTrait
	AutoRepairTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAutoRepair() (x AutoRepair) {
	x.Type = "http://schema.org/AutoRepair"

	return
}
