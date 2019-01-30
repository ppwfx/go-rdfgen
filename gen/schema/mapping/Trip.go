package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTrip strings.Replacer
var strconvTrip strconv.NumError

var basicTripTraitMapping = map[string]func(*schema.TripTrait, []string){}
var customTripTraitMapping = map[string]func(*schema.TripTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Trip", func(ctx jsonld.Context) (interface{}) {
		return NewTripFromContext(ctx)
	})









	customTripTraitMapping["http://schema.org/hasPart"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.HasPart, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.HasPart = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.HasPart = s
		}
	}

	customTripTraitMapping["http://schema.org/isPartOf"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.IsPartOf, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.IsPartOf = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.IsPartOf = s
		}
	}

	customTripTraitMapping["http://schema.org/offers"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		var y schema.Offer
		if strings.HasPrefix(s, "_:") {
			y = NewOfferFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOffer()
			y.Id = s
		}

		x.Offers = &y

		return
	}

	customTripTraitMapping["http://schema.org/provider"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Provider, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Provider = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Provider = s
		}
	}

	customTripTraitMapping["http://schema.org/arrivalTime"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.ArrivalTime = &y

		return
	}

	customTripTraitMapping["http://schema.org/departureTime"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.DepartureTime = &y

		return
	}

	customTripTraitMapping["http://schema.org/itinerary"] = func(x *schema.TripTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Itinerary, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Itinerary = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Itinerary = s
		}
	}
}

func NewTripFromContext(ctx jsonld.Context) (x schema.Trip) {
	x.Type = "http://schema.org/Trip"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTripTrait(ctx, &x.TripTrait)
	MapCustomToTripTrait(ctx, &x.TripTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTripTrait(ctx jsonld.Context, x *schema.TripTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTripTrait(ctx jsonld.Context, x *schema.TripTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTripTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}