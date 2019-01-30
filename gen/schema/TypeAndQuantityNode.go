package schema

type TypeAndQuantityNodeTrait struct {

	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	//
	// RangeIncludes:
	// - http://schema.org/BusinessFunction
	//
	BusinessFunction *BusinessFunction `json:"businessFunction,omitempty" jsonld:"http://schema.org/businessFunction"`

	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	UnitCode interface{} `json:"unitCode,omitempty" jsonld:"http://schema.org/unitCode"`

	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for\n<a href='unitCode'>unitCode</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	UnitText string `json:"unitText,omitempty" jsonld:"http://schema.org/unitText"`

	// The quantity of the goods included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	AmountOfThisGood float64 `json:"amountOfThisGood,omitempty" jsonld:"http://schema.org/amountOfThisGood"`

	// The product that this structured value is referring to.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Service
	//
	TypeOfGood interface{} `json:"typeOfGood,omitempty" jsonld:"http://schema.org/typeOfGood"`

}

type TypeAndQuantityNode struct {
	MetaTrait
	TypeAndQuantityNodeTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTypeAndQuantityNode() (x TypeAndQuantityNode) {
	x.Type = "http://schema.org/TypeAndQuantityNode"

	return
}
