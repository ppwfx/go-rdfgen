package schema

type MusicPlaylistTrait struct {

	// The number of tracks in this album or playlist.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumTracks *Integer `json:"numTracks,omitempty" jsonld:"http://schema.org/numTracks"`

	// A music recording (track)&#x2014;usually a single song. If an ItemList is given, the list should contain items of type MusicRecording.
	//
	// RangeIncludes:
	// - http://schema.org/ItemList
	// - http://schema.org/MusicRecording
	//
	Track interface{} `json:"track,omitempty" jsonld:"http://schema.org/track"`

	// A music recording (track)&#x2014;usually a single song.
	//
	// RangeIncludes:
	// - http://schema.org/MusicRecording
	//
	Tracks *MusicRecording `json:"tracks,omitempty" jsonld:"http://schema.org/tracks"`

}

type MusicPlaylist struct {
	MetaTrait
	MusicPlaylistTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicPlaylist() (x MusicPlaylist) {
	x.Type = "http://schema.org/MusicPlaylist"

	return
}
