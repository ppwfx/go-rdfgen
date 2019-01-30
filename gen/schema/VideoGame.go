package schema

type VideoGameTrait struct {

	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Actor *Person `json:"actor,omitempty" jsonld:"http://schema.org/actor"`

	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Director *Person `json:"director,omitempty" jsonld:"http://schema.org/director"`

	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Actors *Person `json:"actors,omitempty" jsonld:"http://schema.org/actors"`

	// A director of e.g. tv, radio, movie, video games etc. content. Directors can be associated with individual items or with a series, episode, clip.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	//
	Directors *Person `json:"directors,omitempty" jsonld:"http://schema.org/directors"`

	// The composer of the soundtrack.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/MusicGroup
	//
	MusicBy interface{} `json:"musicBy,omitempty" jsonld:"http://schema.org/musicBy"`

	// The trailer of a movie or tv/radio series, season, episode, etc.
	//
	// RangeIncludes:
	// - http://schema.org/VideoObject
	//
	Trailer *VideoObject `json:"trailer,omitempty" jsonld:"http://schema.org/trailer"`

	// Cheat codes to the game.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	CheatCode *CreativeWork `json:"cheatCode,omitempty" jsonld:"http://schema.org/cheatCode"`

	// The electronic systems used to play <a href=\"http://en.wikipedia.org/wiki/Category:Video_game_platforms\">video games</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/URL
	//
	GamePlatform interface{} `json:"gamePlatform,omitempty" jsonld:"http://schema.org/gamePlatform"`

	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	//
	// RangeIncludes:
	// - http://schema.org/GamePlayMode
	//
	PlayMode *GamePlayMode `json:"playMode,omitempty" jsonld:"http://schema.org/playMode"`

	// The server on which  it is possible to play the game.
	//
	// RangeIncludes:
	// - http://schema.org/GameServer
	//
	GameServer *GameServer `json:"gameServer,omitempty" jsonld:"http://schema.org/gameServer"`

	// Links to tips, tactics, etc.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	GameTip *CreativeWork `json:"gameTip,omitempty" jsonld:"http://schema.org/gameTip"`

}

type VideoGame struct {
	MetaTrait
	VideoGameTrait
	GameTrait
	CreativeWorkTrait
	ThingTrait
	SoftwareApplicationTrait
	AdditionalTrait
}

func NewVideoGame() (x VideoGame) {
	x.Type = "http://schema.org/VideoGame"

	return
}
