package schema

type PaymentCardTrait struct {

	// A floor limit is the amount of money above which credit card transactions must be authorized.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	//
	FloorLimit *MonetaryAmount `json:"floorLimit,omitempty" jsonld:"http://schema.org/floorLimit"`

	// A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	// - http://schema.org/Number
	//
	CashBack interface{} `json:"cashBack,omitempty" jsonld:"http://schema.org/cashBack"`

	// A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	ContactlessPayment bool `json:"contactlessPayment,omitempty" jsonld:"http://schema.org/contactlessPayment"`

}

type PaymentCard struct {
	MetaTrait
	PaymentCardTrait
	PaymentMethodTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	FinancialProductTrait
	ServiceTrait
	AdditionalTrait
}

func NewPaymentCard() (x PaymentCard) {
	x.Type = "http://schema.org/PaymentCard"

	return
}
