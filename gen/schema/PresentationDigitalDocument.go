package schema

type PresentationDigitalDocumentTrait struct {

}

type PresentationDigitalDocument struct {
	MetaTrait
	PresentationDigitalDocumentTrait
	DigitalDocumentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewPresentationDigitalDocument() (x PresentationDigitalDocument) {
	x.Type = "http://schema.org/PresentationDigitalDocument"

	return
}
