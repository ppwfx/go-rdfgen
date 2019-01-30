package schema

type MusicAlbumTrait struct {

	// The artist that performed this album or recording.
	//
	// RangeIncludes:
	// - http://schema.org/MusicGroup
	//
	ByArtist *MusicGroup `json:"byArtist,omitempty" jsonld:"http://schema.org/byArtist"`

	// Classification of the album by it's type of content: soundtrack, live album, studio album, etc.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbumProductionType
	//
	AlbumProductionType *MusicAlbumProductionType `json:"albumProductionType,omitempty" jsonld:"http://schema.org/albumProductionType"`

	// A release of this album.
	//
	// RangeIncludes:
	// - http://schema.org/MusicRelease
	//
	AlbumRelease *MusicRelease `json:"albumRelease,omitempty" jsonld:"http://schema.org/albumRelease"`

	// The kind of release which this album is: single, EP or album.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbumReleaseType
	//
	AlbumReleaseType *MusicAlbumReleaseType `json:"albumReleaseType,omitempty" jsonld:"http://schema.org/albumReleaseType"`

}

type MusicAlbum struct {
	MetaTrait
	MusicAlbumTrait
	MusicPlaylistTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicAlbum() (x MusicAlbum) {
	x.Type = "http://schema.org/MusicAlbum"

	return
}
