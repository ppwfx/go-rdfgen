package schema

type ReserveActionTrait struct {

}

type ReserveAction struct {
	MetaTrait
	ReserveActionTrait
	PlanActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReserveAction() (x ReserveAction) {
	x.Type = "http://schema.org/ReserveAction"

	return
}
