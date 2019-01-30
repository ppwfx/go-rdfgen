package schema

type TakeActionTrait struct {

}

type TakeAction struct {
	MetaTrait
	TakeActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTakeAction() (x TakeAction) {
	x.Type = "http://schema.org/TakeAction"

	return
}
