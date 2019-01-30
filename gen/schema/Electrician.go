package schema

type ElectricianTrait struct {

}

type Electrician struct {
	MetaTrait
	ElectricianTrait
	HomeAndConstructionBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewElectrician() (x Electrician) {
	x.Type = "http://schema.org/Electrician"

	return
}
