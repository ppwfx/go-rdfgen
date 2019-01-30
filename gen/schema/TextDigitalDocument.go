package schema

type TextDigitalDocumentTrait struct {

}

type TextDigitalDocument struct {
	MetaTrait
	TextDigitalDocumentTrait
	DigitalDocumentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewTextDigitalDocument() (x TextDigitalDocument) {
	x.Type = "http://schema.org/TextDigitalDocument"

	return
}
