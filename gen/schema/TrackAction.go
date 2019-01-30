package schema

type TrackActionTrait struct {

	// A sub property of instrument. The method of delivery.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod,omitempty" jsonld:"http://schema.org/deliveryMethod"`

}

type TrackAction struct {
	MetaTrait
	TrackActionTrait
	FindActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTrackAction() (x TrackAction) {
	x.Type = "http://schema.org/TrackAction"

	return
}
