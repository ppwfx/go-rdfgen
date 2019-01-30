package schema

type LocalBusinessTrait struct {

	// The larger organization that this local business is a branch of, if any. Not to be confused with (anatomical)<a class=\"localLink\" href=\"http://schema.org/branch\">branch</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	BranchOf *Organization `json:"branchOf,omitempty" jsonld:"http://schema.org/branchOf"`

	// The currency accepted.<br/><br/>\n\nUse standard formats: <a href=\"http://en.wikipedia.org/wiki/ISO_4217\">ISO 4217 currency format</a> e.g. \"USD\"; <a href=\"https://en.wikipedia.org/wiki/List_of_cryptocurrencies\">Ticker symbol</a> for cryptocurrencies e.g. \"BTC\"; well known names for <a href=\"https://en.wikipedia.org/wiki/Local_exchange_trading_system\">Local Exchange Tradings Systems</a> (LETS) and other currency types e.g. \"Ithaca HOUR\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CurrenciesAccepted string `json:"currenciesAccepted,omitempty" jsonld:"http://schema.org/currenciesAccepted"`

	// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas ',' separating each day. Day or time ranges are specified using a hyphen '-'.<br/><br/>\n\n<ul>\n<li>Days are specified using the following two-letter combinations: <code>Mo</code>, <code>Tu</code>, <code>We</code>, <code>Th</code>, <code>Fr</code>, <code>Sa</code>, <code>Su</code>.</li>\n<li>Times are specified using 24:00 time. For example, 3pm is specified as <code>15:00</code>. </li>\n<li>Here is an example: <code>&lt;time itemprop=\"openingHours\" datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;</code>.</li>\n<li>If a business is open 7 days a week, then it can be specified as <code>&lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;</code>.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	OpeningHours string `json:"openingHours,omitempty" jsonld:"http://schema.org/openingHours"`

	// Cash, Credit Card, Cryptocurrency, Local Exchange Tradings System, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PaymentAccepted string `json:"paymentAccepted,omitempty" jsonld:"http://schema.org/paymentAccepted"`

	// The price range of the business, for example <code>$$$</code>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PriceRange string `json:"priceRange,omitempty" jsonld:"http://schema.org/priceRange"`

}

type LocalBusiness struct {
	MetaTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLocalBusiness() (x LocalBusiness) {
	x.Type = "http://schema.org/LocalBusiness"

	return
}
