package schema

type AMRadioChannelTrait struct {

}

type AMRadioChannel struct {
	MetaTrait
	AMRadioChannelTrait
	RadioChannelTrait
	BroadcastChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewAMRadioChannel() (x AMRadioChannel) {
	x.Type = "http://schema.org/AMRadioChannel"

	return
}
