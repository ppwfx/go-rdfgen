package schema

type CompleteDataFeedTrait struct {

}

type CompleteDataFeed struct {
	MetaTrait
	CompleteDataFeedTrait
	DataFeedTrait
	DatasetTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCompleteDataFeed() (x CompleteDataFeed) {
	x.Type = "http://schema.org/CompleteDataFeed"

	return
}
