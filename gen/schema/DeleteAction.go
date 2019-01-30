package schema

type DeleteActionTrait struct {

}

type DeleteAction struct {
	MetaTrait
	DeleteActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDeleteAction() (x DeleteAction) {
	x.Type = "http://schema.org/DeleteAction"

	return
}
