package schema

type UserCommentsTrait struct {

	// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Creator interface{} `json:"creator,omitempty" jsonld:"http://schema.org/creator"`

	// The text of the UserComment.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CommentText string `json:"commentText,omitempty" jsonld:"http://schema.org/commentText"`

	// The time at which the UserComment was made.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	// - http://schema.org/Date
	//
	CommentTime interface{} `json:"commentTime,omitempty" jsonld:"http://schema.org/commentTime"`

	// Specifies the CreativeWork associated with the UserComment.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	Discusses *CreativeWork `json:"discusses,omitempty" jsonld:"http://schema.org/discusses"`

	// The URL at which a reply may be posted to the specified UserComment.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	ReplyToUrl string `json:"replyToUrl,omitempty" jsonld:"http://schema.org/replyToUrl"`

}

type UserComments struct {
	MetaTrait
	UserCommentsTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserComments() (x UserComments) {
	x.Type = "http://schema.org/UserComments"

	return
}
