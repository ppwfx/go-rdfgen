package schema

type DanceGroupTrait struct {

}

type DanceGroup struct {
	MetaTrait
	DanceGroupTrait
	PerformingGroupTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewDanceGroup() (x DanceGroup) {
	x.Type = "http://schema.org/DanceGroup"

	return
}
