package schema

type ExerciseActionTrait struct {

	// Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExerciseType string `json:"exerciseType,omitempty" jsonld:"http://schema.org/exerciseType"`

	// A sub property of location. The original location of the object or the agent before the action.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	FromLocation *Place `json:"fromLocation,omitempty" jsonld:"http://schema.org/fromLocation"`

	// A sub property of location. The final location of the object or the agent after the action.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ToLocation *Place `json:"toLocation,omitempty" jsonld:"http://schema.org/toLocation"`

	// A sub property of location. The course where this action was taken.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ExerciseCourse *Place `json:"exerciseCourse,omitempty" jsonld:"http://schema.org/exerciseCourse"`

	// A sub property of location. The course where this action was taken.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	Course *Place `json:"course,omitempty" jsonld:"http://schema.org/course"`

	// A sub property of participant. The opponent on this action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Opponent *Person `json:"opponent,omitempty" jsonld:"http://schema.org/opponent"`

	// A sub property of instrument. The diet used in this action.
	//
	// RangeIncludes:
	// - http://schema.org/Diet
	//
	Diet *Diet `json:"diet,omitempty" jsonld:"http://schema.org/diet"`

	// The distance travelled, e.g. exercising or travelling.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	//
	Distance *Distance `json:"distance,omitempty" jsonld:"http://schema.org/distance"`

	// A sub property of instrument. The exercise plan used on this action.
	//
	// RangeIncludes:
	// - http://schema.org/ExercisePlan
	//
	ExercisePlan *ExercisePlan `json:"exercisePlan,omitempty" jsonld:"http://schema.org/exercisePlan"`

	// A sub property of instrument. The diet used in this action.
	//
	// RangeIncludes:
	// - http://schema.org/Diet
	//
	ExerciseRelatedDiet *Diet `json:"exerciseRelatedDiet,omitempty" jsonld:"http://schema.org/exerciseRelatedDiet"`

	// A sub property of location. The sports activity location where this action occurred.
	//
	// RangeIncludes:
	// - http://schema.org/SportsActivityLocation
	//
	SportsActivityLocation *SportsActivityLocation `json:"sportsActivityLocation,omitempty" jsonld:"http://schema.org/sportsActivityLocation"`

	// A sub property of location. The sports event where this action occurred.
	//
	// RangeIncludes:
	// - http://schema.org/SportsEvent
	//
	SportsEvent *SportsEvent `json:"sportsEvent,omitempty" jsonld:"http://schema.org/sportsEvent"`

	// A sub property of participant. The sports team that participated on this action.
	//
	// RangeIncludes:
	// - http://schema.org/SportsTeam
	//
	SportsTeam *SportsTeam `json:"sportsTeam,omitempty" jsonld:"http://schema.org/sportsTeam"`

}

type ExerciseAction struct {
	MetaTrait
	ExerciseActionTrait
	PlayActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewExerciseAction() (x ExerciseAction) {
	x.Type = "http://schema.org/ExerciseAction"

	return
}
