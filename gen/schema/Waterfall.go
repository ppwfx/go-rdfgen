package schema

type WaterfallTrait struct {

}

type Waterfall struct {
	MetaTrait
	WaterfallTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewWaterfall() (x Waterfall) {
	x.Type = "http://schema.org/Waterfall"

	return
}
