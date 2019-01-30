package schema

type RoomTrait struct {

}

type Room struct {
	MetaTrait
	RoomTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewRoom() (x Room) {
	x.Type = "http://schema.org/Room"

	return
}
