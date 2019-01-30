package schema

type EndorseActionTrait struct {

	// A sub property of participant. The person/organization being supported.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Endorsee interface{} `json:"endorsee,omitempty" jsonld:"http://schema.org/endorsee"`

}

type EndorseAction struct {
	MetaTrait
	EndorseActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewEndorseAction() (x EndorseAction) {
	x.Type = "http://schema.org/EndorseAction"

	return
}
