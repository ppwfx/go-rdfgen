package schema

type PerformActionTrait struct {

	// A sub property of location. The entertainment business where the action occurred.
	//
	// RangeIncludes:
	// - http://schema.org/EntertainmentBusiness
	//
	EntertainmentBusiness *EntertainmentBusiness `json:"entertainmentBusiness,omitempty" jsonld:"http://schema.org/entertainmentBusiness"`

}

type PerformAction struct {
	MetaTrait
	PerformActionTrait
	PlayActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPerformAction() (x PerformAction) {
	x.Type = "http://schema.org/PerformAction"

	return
}
