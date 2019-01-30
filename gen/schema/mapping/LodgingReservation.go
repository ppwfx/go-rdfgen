package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLodgingReservation strings.Replacer
var strconvLodgingReservation strconv.NumError

var basicLodgingReservationTraitMapping = map[string]func(*schema.LodgingReservationTrait, []string){}
var customLodgingReservationTraitMapping = map[string]func(*schema.LodgingReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LodgingReservation", func(ctx jsonld.Context) (interface{}) {
		return NewLodgingReservationFromContext(ctx)
	})




	basicLodgingReservationTraitMapping["http://schema.org/lodgingUnitDescription"] = func(x *schema.LodgingReservationTrait, s []string) {
		x.LodgingUnitDescription = s[0]
	}





	customLodgingReservationTraitMapping["http://schema.org/checkinTime"] = func(x *schema.LodgingReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CheckinTime = &y

		return
	}

	customLodgingReservationTraitMapping["http://schema.org/checkoutTime"] = func(x *schema.LodgingReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.CheckoutTime = &y

		return
	}

	customLodgingReservationTraitMapping["http://schema.org/lodgingUnitType"] = func(x *schema.LodgingReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LodgingUnitType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LodgingUnitType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LodgingUnitType = s
		}
	}

	customLodgingReservationTraitMapping["http://schema.org/numAdults"] = func(x *schema.LodgingReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumAdults, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumAdults = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumAdults = s
		}
	}

	customLodgingReservationTraitMapping["http://schema.org/numChildren"] = func(x *schema.LodgingReservationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NumChildren, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NumChildren = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NumChildren = s
		}
	}
}

func NewLodgingReservationFromContext(ctx jsonld.Context) (x schema.LodgingReservation) {
	x.Type = "http://schema.org/LodgingReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLodgingReservationTrait(ctx, &x.LodgingReservationTrait)
	MapCustomToLodgingReservationTrait(ctx, &x.LodgingReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLodgingReservationTrait(ctx jsonld.Context, x *schema.LodgingReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLodgingReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLodgingReservationTrait(ctx jsonld.Context, x *schema.LodgingReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLodgingReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}