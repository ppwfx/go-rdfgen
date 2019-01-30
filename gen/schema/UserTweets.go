package schema

type UserTweetsTrait struct {

}

type UserTweets struct {
	MetaTrait
	UserTweetsTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserTweets() (x UserTweets) {
	x.Type = "http://schema.org/UserTweets"

	return
}
