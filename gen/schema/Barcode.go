package schema

type BarcodeTrait struct {

}

type Barcode struct {
	MetaTrait
	BarcodeTrait
	ImageObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewBarcode() (x Barcode) {
	x.Type = "http://schema.org/Barcode"

	return
}
