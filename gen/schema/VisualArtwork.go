package schema

type VisualArtworkTrait struct {

	// The height of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Height interface{} `json:"height,omitempty" jsonld:"http://schema.org/height"`

	// The width of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Distance
	// - http://schema.org/QuantitativeValue
	//
	Width interface{} `json:"width,omitempty" jsonld:"http://schema.org/width"`

	// The depth of the item.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Distance
	//
	Depth interface{} `json:"depth,omitempty" jsonld:"http://schema.org/depth"`

	// The number of copies when multiple copies of a piece of artwork are produced - e.g. for a limited edition of 20 prints, 'artEdition' refers to the total number of copies (in this example \"20\").
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Integer
	//
	ArtEdition interface{} `json:"artEdition,omitempty" jsonld:"http://schema.org/artEdition"`

	// The material used. (e.g. Oil, Watercolour, Acrylic, Linoprint, Marble, Cyanotype, Digital, Lithograph, DryPoint, Intaglio, Pastel, Woodcut, Pencil, Mixed Media, etc.)
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ArtMedium interface{} `json:"artMedium,omitempty" jsonld:"http://schema.org/artMedium"`

	// e.g. Painting, Drawing, Sculpture, Print, Photograph, Assemblage, Collage, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Artform interface{} `json:"artform,omitempty" jsonld:"http://schema.org/artform"`

	// The primary artist for a work\n    in a medium other than pencils or digital line art--for example, if the\n    primary artwork is done in watercolors or digital paints.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Artist *Person `json:"artist,omitempty" jsonld:"http://schema.org/artist"`

	// The supporting materials for the artwork, e.g. Canvas, Paper, Wood, Board, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ArtworkSurface interface{} `json:"artworkSurface,omitempty" jsonld:"http://schema.org/artworkSurface"`

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

	// A material used as a surface in some artwork, e.g. Canvas, Paper, Wood, Board, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Surface interface{} `json:"surface,omitempty" jsonld:"http://schema.org/surface"`

}

type VisualArtwork struct {
	MetaTrait
	VisualArtworkTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewVisualArtwork() (x VisualArtwork) {
	x.Type = "http://schema.org/VisualArtwork"

	return
}
