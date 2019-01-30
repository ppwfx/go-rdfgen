package schema

type ReplyActionTrait struct {

	// A sub property of result. The Comment created or sent as a result of this action.
	//
	// RangeIncludes:
	// - http://schema.org/Comment
	//
	ResultComment *Comment `json:"resultComment,omitempty" jsonld:"http://schema.org/resultComment"`

}

type ReplyAction struct {
	MetaTrait
	ReplyActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReplyAction() (x ReplyAction) {
	x.Type = "http://schema.org/ReplyAction"

	return
}
