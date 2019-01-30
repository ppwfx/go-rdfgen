package schema

type UserPageVisitsTrait struct {

}

type UserPageVisits struct {
	MetaTrait
	UserPageVisitsTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserPageVisits() (x UserPageVisits) {
	x.Type = "http://schema.org/UserPageVisits"

	return
}
