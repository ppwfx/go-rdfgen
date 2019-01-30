package schema

type CheckInActionTrait struct {

}

type CheckInAction struct {
	MetaTrait
	CheckInActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCheckInAction() (x CheckInAction) {
	x.Type = "http://schema.org/CheckInAction"

	return
}
