package schema

type EmailMessageTrait struct {

}

type EmailMessage struct {
	MetaTrait
	EmailMessageTrait
	MessageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewEmailMessage() (x EmailMessage) {
	x.Type = "http://schema.org/EmailMessage"

	return
}
