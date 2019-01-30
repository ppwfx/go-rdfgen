package schema

type MotorcycleRepairTrait struct {

}

type MotorcycleRepair struct {
	MetaTrait
	MotorcycleRepairTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMotorcycleRepair() (x MotorcycleRepair) {
	x.Type = "http://schema.org/MotorcycleRepair"

	return
}
