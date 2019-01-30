package schema

type DataFeedItemTrait struct {

	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	// - http://schema.org/DateTime
	//
	DateCreated interface{} `json:"dateCreated,omitempty" jsonld:"http://schema.org/dateCreated"`

	// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
	//
	// RangeIncludes:
	// - http://schema.org/Date
	// - http://schema.org/DateTime
	//
	DateModified interface{} `json:"dateModified,omitempty" jsonld:"http://schema.org/dateModified"`

	// The datetime the item was removed from the DataFeed.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DateDeleted *DateTime `json:"dateDeleted,omitempty" jsonld:"http://schema.org/dateDeleted"`

	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Item *Thing `json:"item,omitempty" jsonld:"http://schema.org/item"`

}

type DataFeedItem struct {
	MetaTrait
	DataFeedItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDataFeedItem() (x DataFeedItem) {
	x.Type = "http://schema.org/DataFeedItem"

	return
}
