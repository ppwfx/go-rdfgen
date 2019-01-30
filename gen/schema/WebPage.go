package schema

type WebPageTrait struct {

	// Indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the 'speakable' property serves to indicate the parts most likely to be generally useful for speech.<br/><br/>\n\nThe <em>speakable</em> property can be repeated an arbitrary number of times, with three kinds of possible 'content-locator' values:<br/><br/>\n\n1.) <em>id-value</em> URL references - uses <em>id-value</em> of an element in the page being annotated. The simplest use of <em>speakable</em> has (potentially relative) URL values, referencing identified sections of the document concerned.<br/><br/>\n\n2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the <a class=\"localLink\" href=\"http://schema.org/cssSelector\">cssSelector</a> property.<br/><br/>\n\n3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the <a class=\"localLink\" href=\"http://schema.org/xpath\">xpath</a> property.<br/><br/>\n\nFor more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this\nwe define a supporting type, <a class=\"localLink\" href=\"http://schema.org/SpeakableSpecification\">SpeakableSpecification</a>  which is defined to be a possible value of the <em>speakable</em> property.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/SpeakableSpecification
	//
	Speakable interface{} `json:"speakable,omitempty" jsonld:"http://schema.org/speakable"`

	// Indicates the main image on the page.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	//
	PrimaryImageOfPage *ImageObject `json:"primaryImageOfPage,omitempty" jsonld:"http://schema.org/primaryImageOfPage"`

	// A set of links that can help a user understand and navigate a website hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BreadcrumbList
	//
	Breadcrumb interface{} `json:"breadcrumb,omitempty" jsonld:"http://schema.org/breadcrumb"`

	// Date on which the content on this web page was last reviewed for accuracy and/or completeness.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	LastReviewed *Date `json:"lastReviewed,omitempty" jsonld:"http://schema.org/lastReviewed"`

	// Indicates if this web page element is the main subject of the page.
	//
	// RangeIncludes:
	// - http://schema.org/WebPageElement
	//
	MainContentOfPage *WebPageElement `json:"mainContentOfPage,omitempty" jsonld:"http://schema.org/mainContentOfPage"`

	// A link related to this web page, for example to other related web pages.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	RelatedLink string `json:"relatedLink,omitempty" jsonld:"http://schema.org/relatedLink"`

	// People or organizations that have reviewed the content on this web page for accuracy and/or completeness.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	ReviewedBy interface{} `json:"reviewedBy,omitempty" jsonld:"http://schema.org/reviewedBy"`

	// One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	SignificantLink string `json:"significantLink,omitempty" jsonld:"http://schema.org/significantLink"`

	// The most significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	SignificantLinks string `json:"significantLinks,omitempty" jsonld:"http://schema.org/significantLinks"`

	// One of the domain specialities to which this web page's content applies.
	//
	// RangeIncludes:
	// - http://schema.org/Specialty
	//
	Specialty *Specialty `json:"specialty,omitempty" jsonld:"http://schema.org/specialty"`

}

type WebPage struct {
	MetaTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWebPage() (x WebPage) {
	x.Type = "http://schema.org/WebPage"

	return
}
