package schema

type ArriveActionTrait struct {

}

type ArriveAction struct {
	MetaTrait
	ArriveActionTrait
	MoveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewArriveAction() (x ArriveAction) {
	x.Type = "http://schema.org/ArriveAction"

	return
}
