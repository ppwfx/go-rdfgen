package schema

type SocialMediaPostingTrait struct {

	// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	SharedContent *CreativeWork `json:"sharedContent,omitempty" jsonld:"http://schema.org/sharedContent"`

}

type SocialMediaPosting struct {
	MetaTrait
	SocialMediaPostingTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSocialMediaPosting() (x SocialMediaPosting) {
	x.Type = "http://schema.org/SocialMediaPosting"

	return
}
