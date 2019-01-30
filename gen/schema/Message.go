package schema

type MessageTrait struct {

	// A sub property of participant. The participant who is at the receiving end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Audience
	// - http://schema.org/Organization
	//
	Recipient interface{} `json:"recipient,omitempty" jsonld:"http://schema.org/recipient"`

	// The date/time the message was received if a single recipient exists.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DateReceived *DateTime `json:"dateReceived,omitempty" jsonld:"http://schema.org/dateReceived"`

	// The date/time at which the message has been read by the recipient if a single recipient exists.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DateRead *DateTime `json:"dateRead,omitempty" jsonld:"http://schema.org/dateRead"`

	// The date/time at which the message was sent.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DateSent *DateTime `json:"dateSent,omitempty" jsonld:"http://schema.org/dateSent"`

	// A sub property of recipient. The recipient blind copied on a message.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	BccRecipient interface{} `json:"bccRecipient,omitempty" jsonld:"http://schema.org/bccRecipient"`

	// A sub property of recipient. The recipient copied on a message.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	CcRecipient interface{} `json:"ccRecipient,omitempty" jsonld:"http://schema.org/ccRecipient"`

	// A CreativeWork attached to the message.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	MessageAttachment *CreativeWork `json:"messageAttachment,omitempty" jsonld:"http://schema.org/messageAttachment"`

	// A sub property of participant. The participant who is at the sending end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	// - http://schema.org/Audience
	//
	Sender interface{} `json:"sender,omitempty" jsonld:"http://schema.org/sender"`

	// A sub property of recipient. The recipient who was directly sent the message.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Organization
	// - http://schema.org/Audience
	//
	ToRecipient interface{} `json:"toRecipient,omitempty" jsonld:"http://schema.org/toRecipient"`

}

type Message struct {
	MetaTrait
	MessageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMessage() (x Message) {
	x.Type = "http://schema.org/Message"

	return
}
