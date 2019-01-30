package schema

type BlogPostingTrait struct {

}

type BlogPosting struct {
	MetaTrait
	BlogPostingTrait
	SocialMediaPostingTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewBlogPosting() (x BlogPosting) {
	x.Type = "http://schema.org/BlogPosting"

	return
}
