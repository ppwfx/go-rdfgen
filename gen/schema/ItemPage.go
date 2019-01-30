package schema

type ItemPageTrait struct {

}

type ItemPage struct {
	MetaTrait
	ItemPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewItemPage() (x ItemPage) {
	x.Type = "http://schema.org/ItemPage"

	return
}
