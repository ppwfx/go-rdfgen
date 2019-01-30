package schema

type UserBlocksTrait struct {

}

type UserBlocks struct {
	MetaTrait
	UserBlocksTrait
	UserInteractionTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewUserBlocks() (x UserBlocks) {
	x.Type = "http://schema.org/UserBlocks"

	return
}
