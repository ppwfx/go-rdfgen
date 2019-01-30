package schema

type PhotographActionTrait struct {

}

type PhotographAction struct {
	MetaTrait
	PhotographActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPhotographAction() (x PhotographAction) {
	x.Type = "http://schema.org/PhotographAction"

	return
}
