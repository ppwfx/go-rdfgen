package schema

type NewsArticleTrait struct {

	// A <a href=\"https://en.wikipedia.org/wiki/Dateline\">dateline</a> is a brief piece of text included in news articles that describes where and when the story was written or filed though the date is often omitted. Sometimes only a placename is provided.<br/><br/>\n\nStructured representations of dateline-related information can also be expressed more explicitly using <a class=\"localLink\" href=\"http://schema.org/locationCreated\">locationCreated</a> (which represents where a work was created e.g. where a news report was written).  For location depicted or described in the content, use <a class=\"localLink\" href=\"http://schema.org/contentLocation\">contentLocation</a>.<br/><br/>\n\nDateline summaries are oriented more towards human readers than towards automated processing, and can vary substantially. Some examples: \"BEIRUT, Lebanon, June 2.\", \"Paris, France\", \"December 19, 2017 11:43AM Reporting from Washington\", \"Beijing/Moscow\", \"QUEZON CITY, Philippines\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Dateline string `json:"dateline,omitempty" jsonld:"http://schema.org/dateline"`

	// The number of the column in which the NewsArticle appears in the print edition.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PrintColumn string `json:"printColumn,omitempty" jsonld:"http://schema.org/printColumn"`

	// The edition of the print product in which the NewsArticle appears.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PrintEdition string `json:"printEdition,omitempty" jsonld:"http://schema.org/printEdition"`

	// If this NewsArticle appears in print, this field indicates the name of the page on which the article is found. Please note that this field is intended for the exact page name (e.g. A5, B18).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PrintPage string `json:"printPage,omitempty" jsonld:"http://schema.org/printPage"`

	// If this NewsArticle appears in print, this field indicates the print section in which the article appeared.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PrintSection string `json:"printSection,omitempty" jsonld:"http://schema.org/printSection"`

}

type NewsArticle struct {
	MetaTrait
	NewsArticleTrait
	ArticleTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewNewsArticle() (x NewsArticle) {
	x.Type = "http://schema.org/NewsArticle"

	return
}
