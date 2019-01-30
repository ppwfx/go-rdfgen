package schema

type OwnershipInfoTrait struct {

	// The product that this structured value is referring to.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/Service
	//
	TypeOfGood interface{} `json:"typeOfGood,omitempty" jsonld:"http://schema.org/typeOfGood"`

	// The organization or person from which the product was acquired.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	AcquiredFrom interface{} `json:"acquiredFrom,omitempty" jsonld:"http://schema.org/acquiredFrom"`

	// The date and time of obtaining the product.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	OwnedFrom *DateTime `json:"ownedFrom,omitempty" jsonld:"http://schema.org/ownedFrom"`

	// The date and time of giving up ownership on the product.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	OwnedThrough *DateTime `json:"ownedThrough,omitempty" jsonld:"http://schema.org/ownedThrough"`

}

type OwnershipInfo struct {
	MetaTrait
	OwnershipInfoTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOwnershipInfo() (x OwnershipInfo) {
	x.Type = "http://schema.org/OwnershipInfo"

	return
}
