package schema

type MapCategoryTypeTrait struct {

}

type MapCategoryType struct {
	MetaTrait
	MapCategoryTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMapCategoryType() (x MapCategoryType) {
	x.Type = "http://schema.org/MapCategoryType"

	return
}
