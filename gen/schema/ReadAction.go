package schema

type ReadActionTrait struct {

}

type ReadAction struct {
	MetaTrait
	ReadActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReadAction() (x ReadAction) {
	x.Type = "http://schema.org/ReadAction"

	return
}
