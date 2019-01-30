package schema

type ListItemTrait struct {

	// The position of an item in a series or sequence of items.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	// - http://schema.org/Text
	//
	Position interface{} `json:"position,omitempty" jsonld:"http://schema.org/position"`

	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Item *Thing `json:"item,omitempty" jsonld:"http://schema.org/item"`

	// A link to the ListItem that follows the current one.
	//
	// RangeIncludes:
	// - http://schema.org/ListItem
	//
	NextItem *ListItem `json:"nextItem,omitempty" jsonld:"http://schema.org/nextItem"`

	// A link to the ListItem that preceeds the current one.
	//
	// RangeIncludes:
	// - http://schema.org/ListItem
	//
	PreviousItem *ListItem `json:"previousItem,omitempty" jsonld:"http://schema.org/previousItem"`

}

type ListItem struct {
	MetaTrait
	ListItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewListItem() (x ListItem) {
	x.Type = "http://schema.org/ListItem"

	return
}
