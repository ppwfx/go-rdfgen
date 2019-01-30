package schema

type ExerciseGymTrait struct {

}

type ExerciseGym struct {
	MetaTrait
	ExerciseGymTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewExerciseGym() (x ExerciseGym) {
	x.Type = "http://schema.org/ExerciseGym"

	return
}
