package schema

type OrderTrait struct {

	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Seller interface{} `json:"seller,omitempty" jsonld:"http://schema.org/seller"`

	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Broker interface{} `json:"broker,omitempty" jsonld:"http://schema.org/broker"`

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

	// The offer(s) -- e.g., product, quantity and price combinations -- included in the order.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	AcceptedOffer *Offer `json:"acceptedOffer,omitempty" jsonld:"http://schema.org/acceptedOffer"`

	// The billing address for the order.
	//
	// RangeIncludes:
	// - http://schema.org/PostalAddress
	//
	BillingAddress *PostalAddress `json:"billingAddress,omitempty" jsonld:"http://schema.org/billingAddress"`

	// Any discount applied (to an Order).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Discount interface{} `json:"discount,omitempty" jsonld:"http://schema.org/discount"`

	// Code used to redeem a discount.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DiscountCode string `json:"discountCode,omitempty" jsonld:"http://schema.org/discountCode"`

	// The currency of the discount.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DiscountCurrency string `json:"discountCurrency,omitempty" jsonld:"http://schema.org/discountCurrency"`

	// Was the offer accepted as a gift for someone other than the buyer.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsGift bool `json:"isGift,omitempty" jsonld:"http://schema.org/isGift"`

	// 'merchant' is an out-dated term for 'seller'.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Merchant interface{} `json:"merchant,omitempty" jsonld:"http://schema.org/merchant"`

	// Date order was placed.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	OrderDate *DateTime `json:"orderDate,omitempty" jsonld:"http://schema.org/orderDate"`

	// The delivery of the parcel related to this order or order item.
	//
	// RangeIncludes:
	// - http://schema.org/ParcelDelivery
	//
	OrderDelivery *ParcelDelivery `json:"orderDelivery,omitempty" jsonld:"http://schema.org/orderDelivery"`

	// The identifier of the transaction.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	OrderNumber string `json:"orderNumber,omitempty" jsonld:"http://schema.org/orderNumber"`

	// The current status of the order.
	//
	// RangeIncludes:
	// - http://schema.org/OrderStatus
	//
	OrderStatus *OrderStatus `json:"orderStatus,omitempty" jsonld:"http://schema.org/orderStatus"`

	// The item ordered.
	//
	// RangeIncludes:
	// - http://schema.org/Product
	// - http://schema.org/OrderItem
	//
	OrderedItem interface{} `json:"orderedItem,omitempty" jsonld:"http://schema.org/orderedItem"`

	// The order is being paid as part of the referenced Invoice.
	//
	// RangeIncludes:
	// - http://schema.org/Invoice
	//
	PartOfInvoice *Invoice `json:"partOfInvoice,omitempty" jsonld:"http://schema.org/partOfInvoice"`

	// The URL for sending a payment.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	PaymentUrl string `json:"paymentUrl,omitempty" jsonld:"http://schema.org/paymentUrl"`

}

type Order struct {
	MetaTrait
	OrderTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOrder() (x Order) {
	x.Type = "http://schema.org/Order"

	return
}
