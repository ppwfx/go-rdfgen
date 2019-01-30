package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRoom strings.Replacer
var strconvRoom strconv.NumError

var basicRoomTraitMapping = map[string]func(*schema.RoomTrait, []string){}
var customRoomTraitMapping = map[string]func(*schema.RoomTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Room", func(ctx jsonld.Context) (interface{}) {
		return NewRoomFromContext(ctx)
	})

}

func NewRoomFromContext(ctx jsonld.Context) (x schema.Room) {
	x.Type = "http://schema.org/Room"
	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRoomTrait(ctx, &x.RoomTrait)
	MapCustomToRoomTrait(ctx, &x.RoomTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRoomTrait(ctx jsonld.Context, x *schema.RoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRoomTrait(ctx jsonld.Context, x *schema.RoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}