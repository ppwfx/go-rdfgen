package schema

type MusicReleaseFormatTypeTrait struct {

}

type MusicReleaseFormatType struct {
	MetaTrait
	MusicReleaseFormatTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMusicReleaseFormatType() (x MusicReleaseFormatType) {
	x.Type = "http://schema.org/MusicReleaseFormatType"

	return
}
