package schema

type DiscussionForumPostingTrait struct {

}

type DiscussionForumPosting struct {
	MetaTrait
	DiscussionForumPostingTrait
	SocialMediaPostingTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDiscussionForumPosting() (x DiscussionForumPosting) {
	x.Type = "http://schema.org/DiscussionForumPosting"

	return
}
