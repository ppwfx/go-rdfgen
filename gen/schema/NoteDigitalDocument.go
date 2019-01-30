package schema

type NoteDigitalDocumentTrait struct {

}

type NoteDigitalDocument struct {
	MetaTrait
	NoteDigitalDocumentTrait
	DigitalDocumentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewNoteDigitalDocument() (x NoteDigitalDocument) {
	x.Type = "http://schema.org/NoteDigitalDocument"

	return
}
