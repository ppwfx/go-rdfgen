package schema

type TelevisionChannelTrait struct {

}

type TelevisionChannel struct {
	MetaTrait
	TelevisionChannelTrait
	BroadcastChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTelevisionChannel() (x TelevisionChannel) {
	x.Type = "http://schema.org/TelevisionChannel"

	return
}
