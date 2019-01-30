package schema

type CollectionTrait struct {

}

type Collection struct {
	MetaTrait
	CollectionTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCollection() (x Collection) {
	x.Type = "http://schema.org/Collection"

	return
}
