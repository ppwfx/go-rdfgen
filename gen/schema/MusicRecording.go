package schema

type MusicRecordingTrait struct {

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// The artist that performed this album or recording.
	//
	// RangeIncludes:
	// - http://schema.org/MusicGroup
	//
	ByArtist *MusicGroup `json:"byArtist,omitempty" jsonld:"http://schema.org/byArtist"`

	// The album to which this recording belongs.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbum
	//
	InAlbum *MusicAlbum `json:"inAlbum,omitempty" jsonld:"http://schema.org/inAlbum"`

	// The playlist to which this recording belongs.
	//
	// RangeIncludes:
	// - http://schema.org/MusicPlaylist
	//
	InPlaylist *MusicPlaylist `json:"inPlaylist,omitempty" jsonld:"http://schema.org/inPlaylist"`

	// The International Standard Recording Code for the recording.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IsrcCode string `json:"isrcCode,omitempty" jsonld:"http://schema.org/isrcCode"`

	// The composition this track is a recording of.
	//
	// RangeIncludes:
	// - http://schema.org/MusicComposition
	//
	RecordingOf *MusicComposition `json:"recordingOf,omitempty" jsonld:"http://schema.org/recordingOf"`

}

type MusicRecording struct {
	MetaTrait
	MusicRecordingTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicRecording() (x MusicRecording) {
	x.Type = "http://schema.org/MusicRecording"

	return
}
