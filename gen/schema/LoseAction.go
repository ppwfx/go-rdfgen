package schema

type LoseActionTrait struct {

	// A sub property of participant. The winner of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Winner *Person `json:"winner,omitempty" jsonld:"http://schema.org/winner"`

}

type LoseAction struct {
	MetaTrait
	LoseActionTrait
	AchieveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewLoseAction() (x LoseAction) {
	x.Type = "http://schema.org/LoseAction"

	return
}
