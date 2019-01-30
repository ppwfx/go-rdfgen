package schema

type HowToDirectionTrait struct {

	// A sub-property of instrument. A supply consumed when performing instructions or a direction.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/HowToSupply
	//
	Supply interface{} `json:"supply,omitempty" jsonld:"http://schema.org/supply"`

	// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/HowToTool
	//
	Tool interface{} `json:"tool,omitempty" jsonld:"http://schema.org/tool"`

	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	PerformTime *Duration `json:"performTime,omitempty" jsonld:"http://schema.org/performTime"`

	// The length of time it takes to prepare the items to be used in instructions or a direction, in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	PrepTime *Duration `json:"prepTime,omitempty" jsonld:"http://schema.org/prepTime"`

	// The total time required to perform instructions or a direction (including time to prepare the supplies), in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	TotalTime *Duration `json:"totalTime,omitempty" jsonld:"http://schema.org/totalTime"`

	// A media object representing the circumstances after performing this direction.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/MediaObject
	//
	AfterMedia interface{} `json:"afterMedia,omitempty" jsonld:"http://schema.org/afterMedia"`

	// A media object representing the circumstances before performing this direction.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/MediaObject
	//
	BeforeMedia interface{} `json:"beforeMedia,omitempty" jsonld:"http://schema.org/beforeMedia"`

	// A media object representing the circumstances while performing this direction.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/MediaObject
	//
	DuringMedia interface{} `json:"duringMedia,omitempty" jsonld:"http://schema.org/duringMedia"`

}

type HowToDirection struct {
	MetaTrait
	HowToDirectionTrait
	CreativeWorkTrait
	ThingTrait
	ListItemTrait
	IntangibleTrait
	AdditionalTrait
}

func NewHowToDirection() (x HowToDirection) {
	x.Type = "http://schema.org/HowToDirection"

	return
}
