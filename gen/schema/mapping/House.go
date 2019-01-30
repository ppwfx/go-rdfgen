package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHouse strings.Replacer
var strconvHouse strconv.NumError

var basicHouseTraitMapping = map[string]func(*schema.HouseTrait, []string){}
var customHouseTraitMapping = map[string]func(*schema.HouseTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/House", func(ctx jsonld.Context) (interface{}) {
		return NewHouseFromContext(ctx)
	})



	customHouseTraitMapping["http://schema.org/numberOfRooms"] = func(x *schema.HouseTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumberOfRooms, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumberOfRooms = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumberOfRooms = s
		}
	}
}

func NewHouseFromContext(ctx jsonld.Context) (x schema.House) {
	x.Type = "http://schema.org/House"
	MapBasicToAccommodationTrait(ctx, &x.AccommodationTrait)
	MapCustomToAccommodationTrait(ctx, &x.AccommodationTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHouseTrait(ctx, &x.HouseTrait)
	MapCustomToHouseTrait(ctx, &x.HouseTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHouseTrait(ctx jsonld.Context, x *schema.HouseTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHouseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHouseTrait(ctx jsonld.Context, x *schema.HouseTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHouseTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}