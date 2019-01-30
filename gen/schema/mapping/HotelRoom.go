package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHotelRoom strings.Replacer
var strconvHotelRoom strconv.NumError

var basicHotelRoomTraitMapping = map[string]func(*schema.HotelRoomTrait, []string){}
var customHotelRoomTraitMapping = map[string]func(*schema.HotelRoomTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HotelRoom", func(ctx jsonld.Context) (interface{}) {
		return NewHotelRoomFromContext(ctx)
	})




	customHotelRoomTraitMapping["http://schema.org/bed"] = func(x *schema.HotelRoomTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Bed, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Bed = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Bed = s
		}
	}

	customHotelRoomTraitMapping["http://schema.org/occupancy"] = func(x *schema.HotelRoomTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.Occupancy = &y

		return
	}
}

func NewHotelRoomFromContext(ctx jsonld.Context) (x schema.HotelRoom) {
	x.Type = "http://schema.org/HotelRoom"
	MapBasicToRoomTrait(ctx, &x.RoomTrait)
	MapCustomToRoomTrait(ctx, &x.RoomTrait)

	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHotelRoomTrait(ctx, &x.HotelRoomTrait)
	MapCustomToHotelRoomTrait(ctx, &x.HotelRoomTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHotelRoomTrait(ctx jsonld.Context, x *schema.HotelRoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHotelRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHotelRoomTrait(ctx jsonld.Context, x *schema.HotelRoomTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHotelRoomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}