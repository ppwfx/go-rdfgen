package schema

type CategoryCodeSetTrait struct {

	// A Category code contained in this code set.
	//
	// RangeIncludes:
	// - http://schema.org/CategoryCode
	//
	HasCategoryCode *CategoryCode `json:"hasCategoryCode,omitempty" jsonld:"http://schema.org/hasCategoryCode"`

}

type CategoryCodeSet struct {
	MetaTrait
	CategoryCodeSetTrait
	DefinedTermSetTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCategoryCodeSet() (x CategoryCodeSet) {
	x.Type = "http://schema.org/CategoryCodeSet"

	return
}
