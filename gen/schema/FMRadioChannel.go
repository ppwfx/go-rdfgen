package schema

type FMRadioChannelTrait struct {

}

type FMRadioChannel struct {
	MetaTrait
	FMRadioChannelTrait
	RadioChannelTrait
	BroadcastChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFMRadioChannel() (x FMRadioChannel) {
	x.Type = "http://schema.org/FMRadioChannel"

	return
}
