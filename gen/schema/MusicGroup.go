package schema

type MusicGroupTrait struct {

	// Genre of the creative work, broadcast channel or group.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Genre interface{} `json:"genre,omitempty" jsonld:"http://schema.org/genre"`

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

	// A music album.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbum
	//
	Album *MusicAlbum `json:"album,omitempty" jsonld:"http://schema.org/album"`

	// A collection of music albums.
	//
	// RangeIncludes:
	// - http://schema.org/MusicAlbum
	//
	Albums *MusicAlbum `json:"albums,omitempty" jsonld:"http://schema.org/albums"`

	// A member of a music group&#x2014;for example, John, Paul, George, or Ringo.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	MusicGroupMember *Person `json:"musicGroupMember,omitempty" jsonld:"http://schema.org/musicGroupMember"`

}

type MusicGroup struct {
	MetaTrait
	MusicGroupTrait
	PerformingGroupTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicGroup() (x MusicGroup) {
	x.Type = "http://schema.org/MusicGroup"

	return
}
