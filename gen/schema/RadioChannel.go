package schema

type RadioChannelTrait struct {

}

type RadioChannel struct {
	MetaTrait
	RadioChannelTrait
	BroadcastChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRadioChannel() (x RadioChannel) {
	x.Type = "http://schema.org/RadioChannel"

	return
}
