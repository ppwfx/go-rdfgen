package schema

type WarrantyPromiseTrait struct {

	// The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	DurationOfWarranty *QuantitativeValue `json:"durationOfWarranty,omitempty" jsonld:"http://schema.org/durationOfWarranty"`

	// The scope of the warranty promise.
	//
	// RangeIncludes:
	// - http://schema.org/WarrantyScope
	//
	WarrantyScope *WarrantyScope `json:"warrantyScope,omitempty" jsonld:"http://schema.org/warrantyScope"`

}

type WarrantyPromise struct {
	MetaTrait
	WarrantyPromiseTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewWarrantyPromise() (x WarrantyPromise) {
	x.Type = "http://schema.org/WarrantyPromise"

	return
}
