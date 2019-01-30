package schema

type PlanActionTrait struct {

	// The time the object is scheduled to.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ScheduledTime *DateTime `json:"scheduledTime,omitempty" jsonld:"http://schema.org/scheduledTime"`

}

type PlanAction struct {
	MetaTrait
	PlanActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPlanAction() (x PlanAction) {
	x.Type = "http://schema.org/PlanAction"

	return
}
