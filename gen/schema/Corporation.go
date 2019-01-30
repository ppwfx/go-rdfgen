package schema

type CorporationTrait struct {

	// The exchange traded instrument associated with a Corporation object. The tickerSymbol is expressed as an exchange and an instrument name separated by a space character. For the exchange component of the tickerSymbol attribute, we recommend using the controlled vocabulary of Market Identifier Codes (MIC) specified in ISO15022.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TickerSymbol string `json:"tickerSymbol,omitempty" jsonld:"http://schema.org/tickerSymbol"`

}

type Corporation struct {
	MetaTrait
	CorporationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewCorporation() (x Corporation) {
	x.Type = "http://schema.org/Corporation"

	return
}
