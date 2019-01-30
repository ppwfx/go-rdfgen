package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusTrip strings.Replacer
var strconvBusTrip strconv.NumError

var basicBusTripTraitMapping = map[string]func(*schema.BusTripTrait, []string){}
var customBusTripTraitMapping = map[string]func(*schema.BusTripTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusTrip", func(ctx jsonld.Context) (interface{}) {
		return NewBusTripFromContext(ctx)
	})


	basicBusTripTraitMapping["http://schema.org/busNumber"] = func(x *schema.BusTripTrait, s []string) {
		x.BusNumber = s[0]
	}


	basicBusTripTraitMapping["http://schema.org/busName"] = func(x *schema.BusTripTrait, s []string) {
		x.BusName = s[0]
	}




	customBusTripTraitMapping["http://schema.org/arrivalBusStop"] = func(x *schema.BusTripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.ArrivalBusStop, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.ArrivalBusStop = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.ArrivalBusStop = s
		}
	}

	customBusTripTraitMapping["http://schema.org/departureBusStop"] = func(x *schema.BusTripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DepartureBusStop, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DepartureBusStop = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DepartureBusStop = s
		}
	}
}

func NewBusTripFromContext(ctx jsonld.Context) (x schema.BusTrip) {
	x.Type = "http://schema.org/BusTrip"
	MapBasicToTripTrait(ctx, &x.TripTrait)
	MapCustomToTripTrait(ctx, &x.TripTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusTripTrait(ctx, &x.BusTripTrait)
	MapCustomToBusTripTrait(ctx, &x.BusTripTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusTripTrait(ctx jsonld.Context, x *schema.BusTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusTripTrait(ctx jsonld.Context, x *schema.BusTripTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}