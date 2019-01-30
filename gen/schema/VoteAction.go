package schema

type VoteActionTrait struct {

	// A sub property of object. The candidate subject of this action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Candidate *Person `json:"candidate,omitempty" jsonld:"http://schema.org/candidate"`

}

type VoteAction struct {
	MetaTrait
	VoteActionTrait
	ChooseActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewVoteAction() (x VoteAction) {
	x.Type = "http://schema.org/VoteAction"

	return
}
