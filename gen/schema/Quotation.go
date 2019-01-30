package schema

type QuotationTrait struct {

	// The (e.g. fictional) character, Person or Organization to whom the quotation is attributed within the containing CreativeWork.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	SpokenByCharacter interface{} `json:"spokenByCharacter,omitempty" jsonld:"http://schema.org/spokenByCharacter"`

}

type Quotation struct {
	MetaTrait
	QuotationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewQuotation() (x Quotation) {
	x.Type = "http://schema.org/Quotation"

	return
}
