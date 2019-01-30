package schema

type WearActionTrait struct {

}

type WearAction struct {
	MetaTrait
	WearActionTrait
	UseActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewWearAction() (x WearAction) {
	x.Type = "http://schema.org/WearAction"

	return
}
