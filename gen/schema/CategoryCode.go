package schema

type CategoryCodeTrait struct {

	// A short textual code that uniquely identifies the value.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CodeValue string `json:"codeValue,omitempty" jsonld:"http://schema.org/codeValue"`

	// A <a class=\"localLink\" href=\"http://schema.org/CategoryCodeSet\">CategoryCodeSet</a> that contains this catagory code.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/CategoryCodeSet
	//
	InCodeSet interface{} `json:"inCodeSet,omitempty" jsonld:"http://schema.org/inCodeSet"`

}

type CategoryCode struct {
	MetaTrait
	CategoryCodeTrait
	DefinedTermTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewCategoryCode() (x CategoryCode) {
	x.Type = "http://schema.org/CategoryCode"

	return
}
