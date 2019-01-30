package schema

type GasStationTrait struct {

}

type GasStation struct {
	MetaTrait
	GasStationTrait
	AutomotiveBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGasStation() (x GasStation) {
	x.Type = "http://schema.org/GasStation"

	return
}
