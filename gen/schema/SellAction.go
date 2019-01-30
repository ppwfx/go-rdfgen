package schema

type SellActionTrait struct {

	// The warranty promise(s) included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/WarrantyPromise
	//
	WarrantyPromise *WarrantyPromise `json:"warrantyPromise,omitempty" jsonld:"http://schema.org/warrantyPromise"`

	// A sub property of participant. The participant/person/organization that bought the object.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Buyer *Person `json:"buyer,omitempty" jsonld:"http://schema.org/buyer"`

}

type SellAction struct {
	MetaTrait
	SellActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewSellAction() (x SellAction) {
	x.Type = "http://schema.org/SellAction"

	return
}
