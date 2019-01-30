package schema

type GolfCourseTrait struct {

}

type GolfCourse struct {
	MetaTrait
	GolfCourseTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGolfCourse() (x GolfCourse) {
	x.Type = "http://schema.org/GolfCourse"

	return
}
