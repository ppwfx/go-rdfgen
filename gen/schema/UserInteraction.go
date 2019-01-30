package schema

type UserInteractionTrait struct {

}

type UserInteraction struct {
	MetaTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserInteraction() (x UserInteraction) {
	x.Type = "http://schema.org/UserInteraction"

	return
}
