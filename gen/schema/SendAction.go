package schema

type SendActionTrait struct {

	// A sub property of participant. The participant who is at the receiving end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Audience
	// - http://schema.org/Organization
	//
	Recipient interface{} `json:"recipient,omitempty" jsonld:"http://schema.org/recipient"`

	// A sub property of instrument. The method of delivery.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty" jsonld:"http://schema.org/deliveryMethod"`

}

type SendAction struct {
	MetaTrait
	SendActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewSendAction() (x SendAction) {
	x.Type = "http://schema.org/SendAction"

	return
}
