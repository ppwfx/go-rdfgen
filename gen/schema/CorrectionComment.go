package schema

type CorrectionCommentTrait struct {

}

type CorrectionComment struct {
	MetaTrait
	CorrectionCommentTrait
	CommentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCorrectionComment() (x CorrectionComment) {
	x.Type = "http://schema.org/CorrectionComment"

	return
}
