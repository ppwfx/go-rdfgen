package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMeetingRoom strings.Replacer
var strconvMeetingRoom strconv.NumError

var basicMeetingRoomTraitMapping = map[string]func(*schema.MeetingRoomTrait, []string){}
var customMeetingRoomTraitMapping = map[string]func(*schema.MeetingRoomTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MeetingRoom", func(ctx jsonld.Context) (interface{}) {
		return NewMeetingRoomFromContext(ctx)
	})

}

func NewMeetingRoomFromContext(ctx jsonld.Context) (x schema.MeetingRoom) {
	x.Type = "http://schema.org/MeetingRoom"
	MapBasicToRoomTrait(ctx, &x.RoomTrait)
	MapCustomToRoomTrait(ctx, &x.RoomTrait)

	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMeetingRoomTrait(ctx, &x.MeetingRoomTrait)
	MapCustomToMeetingRoomTrait(ctx, &x.MeetingRoomTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMeetingRoomTrait(ctx jsonld.Context, x *schema.MeetingRoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMeetingRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMeetingRoomTrait(ctx jsonld.Context, x *schema.MeetingRoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMeetingRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}