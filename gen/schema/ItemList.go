package schema

type ItemListTrait struct {

	// For itemListElement values, you can use simple strings (e.g. \"Peter\", \"Paul\", \"Mary\"), existing entities, or use ListItem.<br/><br/>\n\nText values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.<br/><br/>\n\nNote: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a 'position' property in such cases.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/ListItem
	//
	ItemListElement interface{} `json:"itemListElement,omitempty" jsonld:"http://schema.org/itemListElement"`

	// Type of ordering (e.g. Ascending, Descending, Unordered).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/ItemListOrderType
	//
	ItemListOrder interface{} `json:"itemListOrder,omitempty" jsonld:"http://schema.org/itemListOrder"`

	// The number of items in an ItemList. Note that some descriptions might not fully describe all items in a list (e.g., multi-page pagination); in such cases, the numberOfItems would be for the entire list.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumberOfItems *Integer `json:"numberOfItems,omitempty" jsonld:"http://schema.org/numberOfItems"`

}

type ItemList struct {
	MetaTrait
	ItemListTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewItemList() (x ItemList) {
	x.Type = "http://schema.org/ItemList"

	return
}
