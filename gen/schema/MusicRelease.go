package schema

type MusicReleaseTrait struct {

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// The catalog number for the release.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CatalogNumber string `json:"catalogNumber,omitempty" jsonld:"http://schema.org/catalogNumber"`

	// The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to \"Stefani Germanotta Band\", but by Lady Gaga.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	CreditedTo interface{} `json:"creditedTo,omitempty" jsonld:"http://schema.org/creditedTo"`

	// Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/MusicReleaseFormatType
	//
	MusicReleaseFormat *MusicReleaseFormatType `json:"musicReleaseFormat,omitempty" jsonld:"http://schema.org/musicReleaseFormat"`

	// The label that issued the release.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	RecordLabel *Organization `json:"recordLabel,omitempty" jsonld:"http://schema.org/recordLabel"`

	// The album this is a release of.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbum
	//
	ReleaseOf *MusicAlbum `json:"releaseOf,omitempty" jsonld:"http://schema.org/releaseOf"`

}

type MusicRelease struct {
	MetaTrait
	MusicReleaseTrait
	MusicPlaylistTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicRelease() (x MusicRelease) {
	x.Type = "http://schema.org/MusicRelease"

	return
}
