package schema

type TVClipTrait struct {

	// The TV series to which this episode or season belongs.
	//
	// RangeIncludes:
	// - http://schema.org/TVSeries
	//
	PartOfTVSeries *TVSeries `json:"partOfTVSeries,omitempty" jsonld:"http://schema.org/partOfTVSeries"`

}

type TVClip struct {
	MetaTrait
	TVClipTrait
	ClipTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewTVClip() (x TVClip) {
	x.Type = "http://schema.org/TVClip"

	return
}
