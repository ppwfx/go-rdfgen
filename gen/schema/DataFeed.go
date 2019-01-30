package schema

type DataFeedTrait struct {

	// An item within in a data feed. Data feeds may have many elements.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/DataFeedItem
	//
	DataFeedElement interface{} `json:"dataFeedElement,omitempty" jsonld:"http://schema.org/dataFeedElement"`

}

type DataFeed struct {
	MetaTrait
	DataFeedTrait
	DatasetTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDataFeed() (x DataFeed) {
	x.Type = "http://schema.org/DataFeed"

	return
}
