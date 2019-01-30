package schema

type MapTrait struct {

	// Indicates the kind of Map, from the MapCategoryType Enumeration.
	//
	// RangeIncludes:
	// - http://schema.org/MapCategoryType
	//
	MapType *MapCategoryType `json:"mapType,omitempty" jsonld:"http://schema.org/mapType"`

}

type Map struct {
	MetaTrait
	MapTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMap() (x Map) {
	x.Type = "http://schema.org/Map"

	return
}
