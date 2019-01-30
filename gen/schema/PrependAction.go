package schema

type PrependActionTrait struct {

}

type PrependAction struct {
	MetaTrait
	PrependActionTrait
	InsertActionTrait
	AddActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPrependAction() (x PrependAction) {
	x.Type = "http://schema.org/PrependAction"

	return
}
