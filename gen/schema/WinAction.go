package schema

type WinActionTrait struct {

	// A sub property of participant. The loser of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Loser *Person `json:"loser,omitempty" jsonld:"http://schema.org/loser"`

}

type WinAction struct {
	MetaTrait
	WinActionTrait
	AchieveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewWinAction() (x WinAction) {
	x.Type = "http://schema.org/WinAction"

	return
}
