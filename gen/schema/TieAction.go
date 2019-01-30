package schema

type TieActionTrait struct {

}

type TieAction struct {
	MetaTrait
	TieActionTrait
	AchieveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewTieAction() (x TieAction) {
	x.Type = "http://schema.org/TieAction"

	return
}
