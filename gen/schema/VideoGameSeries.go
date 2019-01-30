package schema

type VideoGameSeriesTrait struct {

	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	ProductionCompany *Organization `json:"productionCompany,omitempty" jsonld:"http://schema.org/productionCompany"`

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

	// A season that is part of the media series.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeason
	//
	ContainsSeason *CreativeWorkSeason `json:"containsSeason,omitempty" jsonld:"http://schema.org/containsSeason"`

	// An episode of a tv, radio or game media within a series or season.
	//
	// RangeIncludes:
	// - http://schema.org/Episode
	//
	Episode *Episode `json:"episode,omitempty" jsonld:"http://schema.org/episode"`

	// An episode of a TV/radio series or season.
	//
	// RangeIncludes:
	// - http://schema.org/Episode
	//
	Episodes *Episode `json:"episodes,omitempty" jsonld:"http://schema.org/episodes"`

	// The number of episodes in this season or series.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumberOfEpisodes *Integer `json:"numberOfEpisodes,omitempty" jsonld:"http://schema.org/numberOfEpisodes"`

	// The number of seasons in this series.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	NumberOfSeasons *Integer `json:"numberOfSeasons,omitempty" jsonld:"http://schema.org/numberOfSeasons"`

	// A season in a media series.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeason
	//
	Season *CreativeWorkSeason `json:"season,omitempty" jsonld:"http://schema.org/season"`

	// A season in a media series.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWorkSeason
	//
	Seasons *CreativeWorkSeason `json:"seasons,omitempty" jsonld:"http://schema.org/seasons"`

	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	CharacterAttribute *Thing `json:"characterAttribute,omitempty" jsonld:"http://schema.org/characterAttribute"`

	// Cheat codes to the game.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	//
	CheatCode *CreativeWork `json:"cheatCode,omitempty" jsonld:"http://schema.org/cheatCode"`

	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	GameItem *Thing `json:"gameItem,omitempty" jsonld:"http://schema.org/gameItem"`

	// Real or fictional location of the game (or part of game).
	//
	// RangeIncludes:
	// - http://schema.org/URL
	// - http://schema.org/Place
	// - http://schema.org/PostalAddress
	//
	GameLocation interface{} `json:"gameLocation,omitempty" jsonld:"http://schema.org/gameLocation"`

	// The electronic systems used to play <a href=\"http://en.wikipedia.org/wiki/Category:Video_game_platforms\">video games</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/URL
	//
	GamePlatform interface{} `json:"gamePlatform,omitempty" jsonld:"http://schema.org/gamePlatform"`

	// Indicate how many people can play this game (minimum, maximum, or range).
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers,omitempty" jsonld:"http://schema.org/numberOfPlayers"`

	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	//
	// RangeIncludes:
	// - http://schema.org/GamePlayMode
	//
	PlayMode *GamePlayMode `json:"playMode,omitempty" jsonld:"http://schema.org/playMode"`

	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Quest *Thing `json:"quest,omitempty" jsonld:"http://schema.org/quest"`

}

type VideoGameSeries struct {
	MetaTrait
	VideoGameSeriesTrait
	CreativeWorkSeriesTrait
	CreativeWorkTrait
	ThingTrait
	SeriesTrait
	IntangibleTrait
	AdditionalTrait
}

func NewVideoGameSeries() (x VideoGameSeries) {
	x.Type = "http://schema.org/VideoGameSeries"

	return
}
