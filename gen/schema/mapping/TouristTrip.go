package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTouristTrip strings.Replacer
var strconvTouristTrip strconv.NumError

var basicTouristTripTraitMapping = map[string]func(*schema.TouristTripTrait, []string){}
var customTouristTripTraitMapping = map[string]func(*schema.TouristTripTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TouristTrip", func(ctx jsonld.Context) (interface{}) {
		return NewTouristTripFromContext(ctx)
	})



	customTouristTripTraitMapping["http://schema.org/touristType"] = func(x *schema.TouristTripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.TouristType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.TouristType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.TouristType = s
		}
	}
}

func NewTouristTripFromContext(ctx jsonld.Context) (x schema.TouristTrip) {
	x.Type = "http://schema.org/TouristTrip"
	MapBasicToTripTrait(ctx, &x.TripTrait)
	MapCustomToTripTrait(ctx, &x.TripTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTouristTripTrait(ctx, &x.TouristTripTrait)
	MapCustomToTouristTripTrait(ctx, &x.TouristTripTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTouristTripTrait(ctx jsonld.Context, x *schema.TouristTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTouristTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTouristTripTrait(ctx jsonld.Context, x *schema.TouristTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTouristTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}