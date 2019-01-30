package schema

type ResumeActionTrait struct {

}

type ResumeAction struct {
	MetaTrait
	ResumeActionTrait
	ControlActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewResumeAction() (x ResumeAction) {
	x.Type = "http://schema.org/ResumeAction"

	return
}
