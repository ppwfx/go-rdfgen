package schema

type AskActionTrait struct {

	// A sub property of object. A question.
	//
	// RangeIncludes:
	// - http://schema.org/Question
	//
	Question *Question `json:"question,omitempty" jsonld:"http://schema.org/question"`

}

type AskAction struct {
	MetaTrait
	AskActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAskAction() (x AskAction) {
	x.Type = "http://schema.org/AskAction"

	return
}
