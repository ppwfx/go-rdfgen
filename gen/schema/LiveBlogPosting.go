package schema

type LiveBlogPostingTrait struct {

	// The time when the live blog will begin covering the Event. Note that coverage may begin before the Event's start time. The LiveBlogPosting may also be created before coverage begins.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CoverageStartTime *DateTime `json:"coverageStartTime,omitempty" jsonld:"http://schema.org/coverageStartTime"`

	// The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CoverageEndTime *DateTime `json:"coverageEndTime,omitempty" jsonld:"http://schema.org/coverageEndTime"`

	// An update to the LiveBlog.
	//
	// RangeIncludes:
	// - http://schema.org/BlogPosting
	//
	LiveBlogUpdate *BlogPosting `json:"liveBlogUpdate,omitempty" jsonld:"http://schema.org/liveBlogUpdate"`

}

type LiveBlogPosting struct {
	MetaTrait
	LiveBlogPostingTrait
	BlogPostingTrait
	SocialMediaPostingTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewLiveBlogPosting() (x LiveBlogPosting) {
	x.Type = "http://schema.org/LiveBlogPosting"

	return
}
