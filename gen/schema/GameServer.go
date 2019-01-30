package schema

type GameServerTrait struct {

	// Video game which is played on this server.
	//
	// RangeIncludes:
	// - http://schema.org/VideoGame
	//
	Game *VideoGame `json:"game,omitempty" jsonld:"http://schema.org/game"`

	// Number of players on the server.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	PlayersOnline *Integer `json:"playersOnline,omitempty" jsonld:"http://schema.org/playersOnline"`

	// Status of a game server.
	//
	// RangeIncludes:
	// - http://schema.org/GameServerStatus
	//
	ServerStatus *GameServerStatus `json:"serverStatus,omitempty" jsonld:"http://schema.org/serverStatus"`

}

type GameServer struct {
	MetaTrait
	GameServerTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGameServer() (x GameServer) {
	x.Type = "http://schema.org/GameServer"

	return
}
