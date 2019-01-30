package schema

type UserPlusOnesTrait struct {

}

type UserPlusOnes struct {
	MetaTrait
	UserPlusOnesTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserPlusOnes() (x UserPlusOnes) {
	x.Type = "http://schema.org/UserPlusOnes"

	return
}
