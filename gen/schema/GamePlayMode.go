package schema

type GamePlayModeTrait struct {

}

type GamePlayMode struct {
	MetaTrait
	GamePlayModeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGamePlayMode() (x GamePlayMode) {
	x.Type = "http://schema.org/GamePlayMode"

	return
}
