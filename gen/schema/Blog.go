package schema

type BlogTrait struct {

	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Issn string `json:"issn,omitempty" jsonld:"http://schema.org/issn"`

	// A posting that is part of this blog.
	//
	// RangeIncludes:
	// - http://schema.org/BlogPosting
	//
	BlogPost *BlogPosting `json:"blogPost,omitempty" jsonld:"http://schema.org/blogPost"`

	// The postings that are part of this blog.
	//
	// RangeIncludes:
	// - http://schema.org/BlogPosting
	//
	BlogPosts *BlogPosting `json:"blogPosts,omitempty" jsonld:"http://schema.org/blogPosts"`

}

type Blog struct {
	MetaTrait
	BlogTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewBlog() (x Blog) {
	x.Type = "http://schema.org/Blog"

	return
}
