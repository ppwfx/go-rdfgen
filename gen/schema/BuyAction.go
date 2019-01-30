package schema

type BuyActionTrait struct {

	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Seller interface{} `json:"seller,omitempty" jsonld:"http://schema.org/seller"`

	// 'vendor' is an earlier term for 'seller'.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Vendor interface{} `json:"vendor,omitempty" jsonld:"http://schema.org/vendor"`

	// The warranty promise(s) included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/WarrantyPromise
	//
	WarrantyPromise *WarrantyPromise `json:"warrantyPromise,omitempty" jsonld:"http://schema.org/warrantyPromise"`

}

type BuyAction struct {
	MetaTrait
	BuyActionTrait
	TradeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewBuyAction() (x BuyAction) {
	x.Type = "http://schema.org/BuyAction"

	return
}
