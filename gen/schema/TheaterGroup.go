package schema

type TheaterGroupTrait struct {

}

type TheaterGroup struct {
	MetaTrait
	TheaterGroupTrait
	PerformingGroupTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewTheaterGroup() (x TheaterGroup) {
	x.Type = "http://schema.org/TheaterGroup"

	return
}
