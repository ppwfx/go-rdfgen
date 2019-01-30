package schema

type ArticleTrait struct {

	// The actual body of the article.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ArticleBody string `json:"articleBody,omitempty" jsonld:"http://schema.org/articleBody"`

	// Articles may belong to one or more 'sections' in a magazine or newspaper, such as Sports, Lifestyle, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ArticleSection string `json:"articleSection,omitempty" jsonld:"http://schema.org/articleSection"`

	// For an <a class=\"localLink\" href=\"http://schema.org/Article\">Article</a>, typically a <a class=\"localLink\" href=\"http://schema.org/NewsArticle\">NewsArticle</a>, the backstory property provides a textual summary giving a brief explanation of why and how an article was created. In a journalistic setting this could include information about reporting process, methods, interviews, data sources, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	//
	Backstory interface{} `json:"backstory,omitempty" jsonld:"http://schema.org/backstory"`

	// The page on which the work ends; for example \"138\" or \"xvi\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	PageEnd interface{} `json:"pageEnd,omitempty" jsonld:"http://schema.org/pageEnd"`

	// The page on which the work starts; for example \"135\" or \"xiii\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	PageStart interface{} `json:"pageStart,omitempty" jsonld:"http://schema.org/pageStart"`

	// Any description of pages that is not separated into pageStart and pageEnd; for example, \"1-6, 9, 55\" or \"10-12, 46-49\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Pagination string `json:"pagination,omitempty" jsonld:"http://schema.org/pagination"`

	// Indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the 'speakable' property serves to indicate the parts most likely to be generally useful for speech.<br/><br/>\n\nThe <em>speakable</em> property can be repeated an arbitrary number of times, with three kinds of possible 'content-locator' values:<br/><br/>\n\n1.) <em>id-value</em> URL references - uses <em>id-value</em> of an element in the page being annotated. The simplest use of <em>speakable</em> has (potentially relative) URL values, referencing identified sections of the document concerned.<br/><br/>\n\n2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the <a class=\"localLink\" href=\"http://schema.org/cssSelector\">cssSelector</a> property.<br/><br/>\n\n3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the <a class=\"localLink\" href=\"http://schema.org/xpath\">xpath</a> property.<br/><br/>\n\nFor more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this\nwe define a supporting type, <a class=\"localLink\" href=\"http://schema.org/SpeakableSpecification\">SpeakableSpecification</a>  which is defined to be a possible value of the <em>speakable</em> property.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/SpeakableSpecification
	//
	Speakable interface{} `json:"speakable,omitempty" jsonld:"http://schema.org/speakable"`

	// The number of words in the text of the Article.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	WordCount *Integer `json:"wordCount,omitempty" jsonld:"http://schema.org/wordCount"`

}

type Article struct {
	MetaTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewArticle() (x Article) {
	x.Type = "http://schema.org/Article"

	return
}
