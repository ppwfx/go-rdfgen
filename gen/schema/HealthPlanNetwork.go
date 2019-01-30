package schema

type HealthPlanNetworkTrait struct {

	// Whether The costs to the patient for services under this network or formulary.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	HealthPlanCostSharing bool `json:"healthPlanCostSharing,omitempty" jsonld:"http://schema.org/healthPlanCostSharing"`

	// Name or unique ID of network. (Networks are often reused across different insurance plans).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanNetworkId string `json:"healthPlanNetworkId,omitempty" jsonld:"http://schema.org/healthPlanNetworkId"`

	// The tier(s) for this network.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanNetworkTier string `json:"healthPlanNetworkTier,omitempty" jsonld:"http://schema.org/healthPlanNetworkTier"`

}

type HealthPlanNetwork struct {
	MetaTrait
	HealthPlanNetworkTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewHealthPlanNetwork() (x HealthPlanNetwork) {
	x.Type = "http://schema.org/HealthPlanNetwork"

	return
}
