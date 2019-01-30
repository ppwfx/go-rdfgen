package schema

type PerformanceRoleTrait struct {

	// The name of a character played in some acting or performing role, i.e. in a PerformanceRole.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CharacterName string `json:"characterName,omitempty" jsonld:"http://schema.org/characterName"`

}

type PerformanceRole struct {
	MetaTrait
	PerformanceRoleTrait
	RoleTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPerformanceRole() (x PerformanceRole) {
	x.Type = "http://schema.org/PerformanceRole"

	return
}
