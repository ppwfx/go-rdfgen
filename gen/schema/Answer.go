package schema

type AnswerTrait struct {

}

type Answer struct {
	MetaTrait
	AnswerTrait
	CommentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewAnswer() (x Answer) {
	x.Type = "http://schema.org/Answer"

	return
}
