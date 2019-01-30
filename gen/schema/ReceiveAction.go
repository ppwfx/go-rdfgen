package schema

type ReceiveActionTrait struct {

	// A sub property of participant. The participant who is at the sending end of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	// - http://schema.org/Audience
	//
	Sender interface{} `json:"sender,omitempty" jsonld:"http://schema.org/sender"`

	// A sub property of instrument. The method of delivery.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty" jsonld:"http://schema.org/deliveryMethod"`

}

type ReceiveAction struct {
	MetaTrait
	ReceiveActionTrait
	TransferActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReceiveAction() (x ReceiveAction) {
	x.Type = "http://schema.org/ReceiveAction"

	return
}
