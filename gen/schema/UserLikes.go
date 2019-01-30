package schema

type UserLikesTrait struct {

}

type UserLikes struct {
	MetaTrait
	UserLikesTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserLikes() (x UserLikes) {
	x.Type = "http://schema.org/UserLikes"

	return
}
