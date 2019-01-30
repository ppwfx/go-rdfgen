package schema

type ExercisePlanTrait struct {

	// How often one should break from the activity.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	RestPeriods interface{} `json:"restPeriods,omitempty" jsonld:"http://schema.org/restPeriods"`

	// How often one should engage in the activity.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	ActivityFrequency interface{} `json:"activityFrequency,omitempty" jsonld:"http://schema.org/activityFrequency"`

	// Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExerciseType string `json:"exerciseType,omitempty" jsonld:"http://schema.org/exerciseType"`

	// Quantitative measure gauging the degree of force involved in the exercise, for example, heartbeats per minute. May include the velocity of the movement.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	Intensity interface{} `json:"intensity,omitempty" jsonld:"http://schema.org/intensity"`

	// Any additional component of the exercise prescription that may need to be articulated to the patient. This may include the order of exercises, the number of repetitions of movement, quantitative distance, progressions over time, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AdditionalVariable string `json:"additionalVariable,omitempty" jsonld:"http://schema.org/additionalVariable"`

	// Length of time to engage in the activity.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	// - http://schema.org/QualitativeValue
	//
	ActivityDuration interface{} `json:"activityDuration,omitempty" jsonld:"http://schema.org/activityDuration"`

	// Number of times one should repeat the activity.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	// - http://schema.org/QualitativeValue
	//
	Repetitions interface{} `json:"repetitions,omitempty" jsonld:"http://schema.org/repetitions"`

	// Quantitative measure of the physiologic output of the exercise; also referred to as energy expenditure.
	//
	// RangeIncludes:
	// - http://schema.org/Energy
	// - http://schema.org/QualitativeValue
	//
	Workload interface{} `json:"workload,omitempty" jsonld:"http://schema.org/workload"`

}

type ExercisePlan struct {
	MetaTrait
	ExercisePlanTrait
	CreativeWorkTrait
	ThingTrait
	PhysicalActivityTrait
	LifestyleModificationTrait
	MedicalEntityTrait
	AdditionalTrait
}

func NewExercisePlan() (x ExercisePlan) {
	x.Type = "http://schema.org/ExercisePlan"

	return
}
