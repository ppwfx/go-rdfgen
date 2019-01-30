package schema

type IndividualProductTrait struct {

	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SerialNumber string `json:"serialNumber,omitempty" jsonld:"http://schema.org/serialNumber"`

}

type IndividualProduct struct {
	MetaTrait
	IndividualProductTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewIndividualProduct() (x IndividualProduct) {
	x.Type = "http://schema.org/IndividualProduct"

	return
}
