package schema

type RecyclingCenterTrait struct {

}

type RecyclingCenter struct {
	MetaTrait
	RecyclingCenterTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewRecyclingCenter() (x RecyclingCenter) {
	x.Type = "http://schema.org/RecyclingCenter"

	return
}
