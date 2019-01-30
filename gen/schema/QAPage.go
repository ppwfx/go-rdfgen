package schema

type QAPageTrait struct {

}

type QAPage struct {
	MetaTrait
	QAPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewQAPage() (x QAPage) {
	x.Type = "http://schema.org/QAPage"

	return
}
