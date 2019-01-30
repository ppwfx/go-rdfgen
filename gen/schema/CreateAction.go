package schema

type CreateActionTrait struct {

}

type CreateAction struct {
	MetaTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCreateAction() (x CreateAction) {
	x.Type = "http://schema.org/CreateAction"

	return
}
