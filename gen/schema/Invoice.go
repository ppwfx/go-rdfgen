package schema

type InvoiceTrait struct {

	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Provider interface{} `json:"provider,omitempty" jsonld:"http://schema.org/provider"`

	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/PhysicalActivityCategory
	//
	Category interface{} `json:"category,omitempty" jsonld:"http://schema.org/category"`

	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Broker interface{} `json:"broker,omitempty" jsonld:"http://schema.org/broker"`

	// The identifier for the account the payment will be applied to.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AccountId string `json:"accountId,omitempty" jsonld:"http://schema.org/accountId"`

	// The time interval used to compute the invoice.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	BillingPeriod *Duration `json:"billingPeriod,omitempty" jsonld:"http://schema.org/billingPeriod"`

	// A number that confirms the given order or payment has been received.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ConfirmationNumber string `json:"confirmationNumber,omitempty" jsonld:"http://schema.org/confirmationNumber"`

	// Party placing the order or paying the invoice.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Customer interface{} `json:"customer,omitempty" jsonld:"http://schema.org/customer"`

	// The minimum payment required at this time.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/PriceSpecification
	//
	MinimumPaymentDue interface{} `json:"minimumPaymentDue,omitempty" jsonld:"http://schema.org/minimumPaymentDue"`

	// The date that payment is due.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	PaymentDue *DateTime `json:"paymentDue,omitempty" jsonld:"http://schema.org/paymentDue"`

	// The date that payment is due.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	PaymentDueDate *DateTime `json:"paymentDueDate,omitempty" jsonld:"http://schema.org/paymentDueDate"`

	// The name of the credit card or other method of payment for the order.
	//
	// RangeIncludes:
	// - http://schema.org/PaymentMethod
	//
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty" jsonld:"http://schema.org/paymentMethod"`

	// An identifier for the method of payment used (e.g. the last 4 digits of the credit card).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PaymentMethodId string `json:"paymentMethodId,omitempty" jsonld:"http://schema.org/paymentMethodId"`

	// The status of payment; whether the invoice has been paid or not.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PaymentStatusType
	//
	PaymentStatus interface{} `json:"paymentStatus,omitempty" jsonld:"http://schema.org/paymentStatus"`

	// The Order(s) related to this Invoice. One or more Orders may be combined into a single Invoice.
	//
	// RangeIncludes:
	// - http://schema.org/Order
	//
	ReferencesOrder *Order `json:"referencesOrder,omitempty" jsonld:"http://schema.org/referencesOrder"`

	// The date the invoice is scheduled to be paid.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	//
	ScheduledPaymentDate *Date `json:"scheduledPaymentDate,omitempty" jsonld:"http://schema.org/scheduledPaymentDate"`

	// The total amount due.
	//
	// RangeIncludes:
	// - http://schema.org/MonetaryAmount
	// - http://schema.org/PriceSpecification
	//
	TotalPaymentDue interface{} `json:"totalPaymentDue,omitempty" jsonld:"http://schema.org/totalPaymentDue"`

}

type Invoice struct {
	MetaTrait
	InvoiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewInvoice() (x Invoice) {
	x.Type = "http://schema.org/Invoice"

	return
}
