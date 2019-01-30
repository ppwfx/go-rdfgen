package schema

type GameServerStatusTrait struct {

}

type GameServerStatus struct {
	MetaTrait
	GameServerStatusTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGameServerStatus() (x GameServerStatus) {
	x.Type = "http://schema.org/GameServerStatus"

	return
}
