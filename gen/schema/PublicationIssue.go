package schema

type PublicationIssueTrait struct {

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

	// Identifies the issue of publication; for example, \"iii\" or \"2\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	IssueNumber interface{} `json:"issueNumber,omitempty" jsonld:"http://schema.org/issueNumber"`

}

type PublicationIssue struct {
	MetaTrait
	PublicationIssueTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewPublicationIssue() (x PublicationIssue) {
	x.Type = "http://schema.org/PublicationIssue"

	return
}
