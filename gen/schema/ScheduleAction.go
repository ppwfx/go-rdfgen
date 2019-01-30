package schema

type ScheduleActionTrait struct {

}

type ScheduleAction struct {
	MetaTrait
	ScheduleActionTrait
	PlanActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewScheduleAction() (x ScheduleAction) {
	x.Type = "http://schema.org/ScheduleAction"

	return
}
