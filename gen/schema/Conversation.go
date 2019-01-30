package schema

type ConversationTrait struct {

}

type Conversation struct {
	MetaTrait
	ConversationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewConversation() (x Conversation) {
	x.Type = "http://schema.org/Conversation"

	return
}
