package schema

type SpreadsheetDigitalDocumentTrait struct {

}

type SpreadsheetDigitalDocument struct {
	MetaTrait
	SpreadsheetDigitalDocumentTrait
	DigitalDocumentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewSpreadsheetDigitalDocument() (x SpreadsheetDigitalDocument) {
	x.Type = "http://schema.org/SpreadsheetDigitalDocument"

	return
}
