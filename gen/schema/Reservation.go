package schema

type ReservationTrait struct {

	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Provider interface{} `json:"provider,omitempty" jsonld:"http://schema.org/provider"`

	// The currency of the price, or a price component when attached to <a class=\"localLink\" href=\"http://schema.org/PriceSpecification\">PriceSpecification</a> and its subtypes.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceCurrency string `json:"priceCurrency,omitempty" jsonld:"http://schema.org/priceCurrency"`

	// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Broker interface{} `json:"broker,omitempty" jsonld:"http://schema.org/broker"`

	// 'bookingAgent' is an out-dated term indicating a 'broker' that serves as a booking agent.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	BookingAgent interface{} `json:"bookingAgent,omitempty" jsonld:"http://schema.org/bookingAgent"`

	// The date and time the reservation was booked.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	BookingTime *DateTime `json:"bookingTime,omitempty" jsonld:"http://schema.org/bookingTime"`

	// The date and time the reservation was modified.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ModifiedTime *DateTime `json:"modifiedTime,omitempty" jsonld:"http://schema.org/modifiedTime"`

	// Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	//
	// RangeIncludes:
	// - http://schema.org/ProgramMembership
	//
	ProgramMembershipUsed *ProgramMembership `json:"programMembershipUsed,omitempty" jsonld:"http://schema.org/programMembershipUsed"`

	// The thing -- flight, event, restaurant,etc. being reserved.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	ReservationFor *Thing `json:"reservationFor,omitempty" jsonld:"http://schema.org/reservationFor"`

	// A unique identifier for the reservation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ReservationId string `json:"reservationId,omitempty" jsonld:"http://schema.org/reservationId"`

	// The current status of the reservation.
	//
	// RangeIncludes:
	// - http://schema.org/ReservationStatusType
	//
	ReservationStatus *ReservationStatusType `json:"reservationStatus,omitempty" jsonld:"http://schema.org/reservationStatus"`

	// A ticket associated with the reservation.
	//
	// RangeIncludes:
	// - http://schema.org/Ticket
	//
	ReservedTicket *Ticket `json:"reservedTicket,omitempty" jsonld:"http://schema.org/reservedTicket"`

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

}

type Reservation struct {
	MetaTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewReservation() (x Reservation) {
	x.Type = "http://schema.org/Reservation"

	return
}
