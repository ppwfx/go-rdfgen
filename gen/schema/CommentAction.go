package schema

type CommentActionTrait struct {

	// A sub property of result. The Comment created or sent as a result of this action.
	//
	// RangeIncludes:
	// - http://schema.org/Comment
	//
	ResultComment *Comment `json:"resultComment,omitempty" jsonld:"http://schema.org/resultComment"`

}

type CommentAction struct {
	MetaTrait
	CommentActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCommentAction() (x CommentAction) {
	x.Type = "http://schema.org/CommentAction"

	return
}
