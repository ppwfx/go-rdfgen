package schema

type FollowActionTrait struct {

	// A sub property of object. The person or organization being followed.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Followee interface{} `json:"followee,omitempty" jsonld:"http://schema.org/followee"`

}

type FollowAction struct {
	MetaTrait
	FollowActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewFollowAction() (x FollowAction) {
	x.Type = "http://schema.org/FollowAction"

	return
}
