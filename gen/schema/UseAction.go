package schema

type UseActionTrait struct {

}

type UseAction struct {
	MetaTrait
	UseActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewUseAction() (x UseAction) {
	x.Type = "http://schema.org/UseAction"

	return
}
