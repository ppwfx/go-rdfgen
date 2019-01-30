package schema

type AchieveActionTrait struct {

}

type AchieveAction struct {
	MetaTrait
	AchieveActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAchieveAction() (x AchieveAction) {
	x.Type = "http://schema.org/AchieveAction"

	return
}
