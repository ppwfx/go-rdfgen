package schema

type UserCheckinsTrait struct {

}

type UserCheckins struct {
	MetaTrait
	UserCheckinsTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserCheckins() (x UserCheckins) {
	x.Type = "http://schema.org/UserCheckins"

	return
}
