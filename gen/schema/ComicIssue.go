package schema

type ComicIssueTrait struct {

	// The primary artist for a work\n    in a medium other than pencils or digital line art--for example, if the\n    primary artwork is done in watercolors or digital paints.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Artist *Person `json:"artist,omitempty" jsonld:"http://schema.org/artist"`

	// The individual who adds color to inked drawings.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Colorist *Person `json:"colorist,omitempty" jsonld:"http://schema.org/colorist"`

	// The individual who traces over the pencil drawings in ink after pencils are complete.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Inker *Person `json:"inker,omitempty" jsonld:"http://schema.org/inker"`

	// The individual who adds lettering, including speech balloons and sound effects, to artwork.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Letterer *Person `json:"letterer,omitempty" jsonld:"http://schema.org/letterer"`

	// The individual who draws the primary narrative artwork.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Penciler *Person `json:"penciler,omitempty" jsonld:"http://schema.org/penciler"`

	// A description of the variant cover\n    for the issue, if the issue is a variant printing. For example, \"Bryan Hitch\n    Variant Cover\" or \"2nd Printing Variant\".
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	VariantCover string `json:"variantCover,omitempty" jsonld:"http://schema.org/variantCover"`

}

type ComicIssue struct {
	MetaTrait
	ComicIssueTrait
	PublicationIssueTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewComicIssue() (x ComicIssue) {
	x.Type = "http://schema.org/ComicIssue"

	return
}
