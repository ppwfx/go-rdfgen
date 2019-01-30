package schema

type MeetingRoomTrait struct {

}

type MeetingRoom struct {
	MetaTrait
	MeetingRoomTrait
	RoomTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewMeetingRoom() (x MeetingRoom) {
	x.Type = "http://schema.org/MeetingRoom"

	return
}
