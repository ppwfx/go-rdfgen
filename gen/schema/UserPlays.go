package schema

type UserPlaysTrait struct {

}

type UserPlays struct {
	MetaTrait
	UserPlaysTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserPlays() (x UserPlays) {
	x.Type = "http://schema.org/UserPlays"

	return
}
