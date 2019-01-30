package schema

type TicketTrait struct {

	// The organization issuing the ticket or permit.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	IssuedBy *Organization `json:"issuedBy,omitempty" jsonld:"http://schema.org/issuedBy"`

	// The currency of the price, or a price component when attached to <a class=\"localLink\" href=\"http://schema.org/PriceSpecification\">PriceSpecification</a> and its subtypes.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceCurrency string `json:"priceCurrency,omitempty" jsonld:"http://schema.org/priceCurrency"`

	// The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PriceSpecification
	// - http://schema.org/Number
	//
	TotalPrice interface{} `json:"totalPrice,omitempty" jsonld:"http://schema.org/totalPrice"`

	// The person or organization the reservation or ticket is for.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	UnderName interface{} `json:"underName,omitempty" jsonld:"http://schema.org/underName"`

	// The date the ticket was issued.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DateIssued *DateTime `json:"dateIssued,omitempty" jsonld:"http://schema.org/dateIssued"`

	// The unique identifier for the ticket.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TicketNumber string `json:"ticketNumber,omitempty" jsonld:"http://schema.org/ticketNumber"`

	// Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	TicketToken interface{} `json:"ticketToken,omitempty" jsonld:"http://schema.org/ticketToken"`

	// The seat associated with the ticket.
	//
	// RangeIncludes:
	// - http://schema.org/Seat
	//
	TicketedSeat *Seat `json:"ticketedSeat,omitempty" jsonld:"http://schema.org/ticketedSeat"`

}

type Ticket struct {
	MetaTrait
	TicketTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTicket() (x Ticket) {
	x.Type = "http://schema.org/Ticket"

	return
}
