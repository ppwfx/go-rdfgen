package schema

type GameTrait struct {

	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	CharacterAttribute *Thing `json:"characterAttribute,omitempty" jsonld:"http://schema.org/characterAttribute"`

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

	// Indicate how many people can play this game (minimum, maximum, or range).
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers,omitempty" jsonld:"http://schema.org/numberOfPlayers"`

	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	Quest *Thing `json:"quest,omitempty" jsonld:"http://schema.org/quest"`

}

type Game struct {
	MetaTrait
	GameTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewGame() (x Game) {
	x.Type = "http://schema.org/Game"

	return
}
