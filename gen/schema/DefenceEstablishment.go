package schema

type DefenceEstablishmentTrait struct {

}

type DefenceEstablishment struct {
	MetaTrait
	DefenceEstablishmentTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewDefenceEstablishment() (x DefenceEstablishment) {
	x.Type = "http://schema.org/DefenceEstablishment"

	return
}
