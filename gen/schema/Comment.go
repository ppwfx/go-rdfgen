package schema

type CommentTrait struct {

	// The number of downvotes this question, answer or comment has received from the community.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	DownvoteCount *Integer `json:"downvoteCount,omitempty" jsonld:"http://schema.org/downvoteCount"`

	// The parent of a question, answer or item in general.
	//
	// RangeIncludes:
	// - http://schema.org/Question
	//
	ParentItem *Question `json:"parentItem,omitempty" jsonld:"http://schema.org/parentItem"`

	// The number of upvotes this question, answer or comment has received from the community.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	UpvoteCount *Integer `json:"upvoteCount,omitempty" jsonld:"http://schema.org/upvoteCount"`

}

type Comment struct {
	MetaTrait
	CommentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewComment() (x Comment) {
	x.Type = "http://schema.org/Comment"

	return
}
