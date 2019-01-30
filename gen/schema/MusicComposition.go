package schema

type MusicCompositionTrait struct {

	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Composer interface{} `json:"composer,omitempty" jsonld:"http://schema.org/composer"`

	// The date and place the work was first performed.
	//
	// RangeIncludes:
	// - http://schema.org/Event
	//
	FirstPerformance *Event `json:"firstPerformance,omitempty" jsonld:"http://schema.org/firstPerformance"`

	// Smaller compositions included in this work (e.g. a movement in a symphony).
	//
	// RangeIncludes:
	// - http://schema.org/MusicComposition
	//
	IncludedComposition *MusicComposition `json:"includedComposition,omitempty" jsonld:"http://schema.org/includedComposition"`

	// The International Standard Musical Work Code for the composition.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IswcCode string `json:"iswcCode,omitempty" jsonld:"http://schema.org/iswcCode"`

	// The person who wrote the words.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Lyricist *Person `json:"lyricist,omitempty" jsonld:"http://schema.org/lyricist"`

	// The words in the song.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	Lyrics *CreativeWork `json:"lyrics,omitempty" jsonld:"http://schema.org/lyrics"`

	// An arrangement derived from the composition.
	//
	// RangeIncludes:
	// - http://schema.org/MusicComposition
	//
	MusicArrangement *MusicComposition `json:"musicArrangement,omitempty" jsonld:"http://schema.org/musicArrangement"`

	// The type of composition (e.g. overture, sonata, symphony, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MusicCompositionForm string `json:"musicCompositionForm,omitempty" jsonld:"http://schema.org/musicCompositionForm"`

	// The key, mode, or scale this composition uses.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MusicalKey string `json:"musicalKey,omitempty" jsonld:"http://schema.org/musicalKey"`

	// An audio recording of the work.
	//
	// RangeIncludes:
	// - http://schema.org/MusicRecording
	//
	RecordedAs *MusicRecording `json:"recordedAs,omitempty" jsonld:"http://schema.org/recordedAs"`

}

type MusicComposition struct {
	MetaTrait
	MusicCompositionTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicComposition() (x MusicComposition) {
	x.Type = "http://schema.org/MusicComposition"

	return
}
