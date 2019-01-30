package schema

type PerformingGroupTrait struct {

}

type PerformingGroup struct {
	MetaTrait
	PerformingGroupTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewPerformingGroup() (x PerformingGroup) {
	x.Type = "http://schema.org/PerformingGroup"

	return
}
