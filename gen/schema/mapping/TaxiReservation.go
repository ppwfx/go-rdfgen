package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTaxiReservation strings.Replacer
var strconvTaxiReservation strconv.NumError

var basicTaxiReservationTraitMapping = map[string]func(*schema.TaxiReservationTrait, []string){}
var customTaxiReservationTraitMapping = map[string]func(*schema.TaxiReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TaxiReservation", func(ctx jsonld.Context) (interface{}) {
		return NewTaxiReservationFromContext(ctx)
	})





	customTaxiReservationTraitMapping["http://schema.org/pickupTime"] = func(x *schema.TaxiReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.PickupTime = &y

		return
	}

	customTaxiReservationTraitMapping["http://schema.org/partySize"] = func(x *schema.TaxiReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.PartySize, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.PartySize = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.PartySize = s
		}
	}

	customTaxiReservationTraitMapping["http://schema.org/pickupLocation"] = func(x *schema.TaxiReservationTrait, ctx jsonld.Context, s string){
		var y schema.Place
		if strings.HasPrefix(s, "_:") {
			y = NewPlaceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPlace()
			y.Id = s
		}

		x.PickupLocation = &y

		return
	}
}

func NewTaxiReservationFromContext(ctx jsonld.Context) (x schema.TaxiReservation) {
	x.Type = "http://schema.org/TaxiReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTaxiReservationTrait(ctx, &x.TaxiReservationTrait)
	MapCustomToTaxiReservationTrait(ctx, &x.TaxiReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTaxiReservationTrait(ctx jsonld.Context, x *schema.TaxiReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTaxiReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTaxiReservationTrait(ctx jsonld.Context, x *schema.TaxiReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTaxiReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}