package schema

type ListenActionTrait struct {

}

type ListenAction struct {
	MetaTrait
	ListenActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewListenAction() (x ListenAction) {
	x.Type = "http://schema.org/ListenAction"

	return
}
