package schema

type CasinoTrait struct {

}

type Casino struct {
	MetaTrait
	CasinoTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewCasino() (x Casino) {
	x.Type = "http://schema.org/Casino"

	return
}
