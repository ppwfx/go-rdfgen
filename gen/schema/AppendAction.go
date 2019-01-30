package schema

type AppendActionTrait struct {

}

type AppendAction struct {
	MetaTrait
	AppendActionTrait
	InsertActionTrait
	AddActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAppendAction() (x AppendAction) {
	x.Type = "http://schema.org/AppendAction"

	return
}
