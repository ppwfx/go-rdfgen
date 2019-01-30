package schema

type PublicSwimmingPoolTrait struct {

}

type PublicSwimmingPool struct {
	MetaTrait
	PublicSwimmingPoolTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPublicSwimmingPool() (x PublicSwimmingPool) {
	x.Type = "http://schema.org/PublicSwimmingPool"

	return
}
