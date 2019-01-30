package schema

type SocialEventTrait struct {

}

type SocialEvent struct {
	MetaTrait
	SocialEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewSocialEvent() (x SocialEvent) {
	x.Type = "http://schema.org/SocialEvent"

	return
}
