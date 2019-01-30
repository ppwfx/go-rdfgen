package schema

type QuestionTrait struct {

	// The number of downvotes this question, answer or comment has received from the community.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	DownvoteCount *Integer `json:"downvoteCount,omitempty" jsonld:"http://schema.org/downvoteCount"`

	// The number of upvotes this question, answer or comment has received from the community.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	UpvoteCount *Integer `json:"upvoteCount,omitempty" jsonld:"http://schema.org/upvoteCount"`

	// The answer(s) that has been accepted as best, typically on a Question/Answer site. Sites vary in their selection mechanisms, e.g. drawing on community opinion and/or the view of the Question author.
	//
	// RangeIncludes:
	// - http://schema.org/ItemList
	// - http://schema.org/Answer
	//
	AcceptedAnswer interface{} `json:"acceptedAnswer,omitempty" jsonld:"http://schema.org/acceptedAnswer"`

	// The number of answers this question has received.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	AnswerCount *Integer `json:"answerCount,omitempty" jsonld:"http://schema.org/answerCount"`

	// An answer (possibly one of several, possibly incorrect) to a Question, e.g. on a Question/Answer site.
	//
	// RangeIncludes:
	// - http://schema.org/ItemList
	// - http://schema.org/Answer
	//
	SuggestedAnswer interface{} `json:"suggestedAnswer,omitempty" jsonld:"http://schema.org/suggestedAnswer"`

}

type Question struct {
	MetaTrait
	QuestionTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewQuestion() (x Question) {
	x.Type = "http://schema.org/Question"

	return
}
