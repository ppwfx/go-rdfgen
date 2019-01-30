package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFoodEstablishmentReservation strings.Replacer
var strconvFoodEstablishmentReservation strconv.NumError

var basicFoodEstablishmentReservationTraitMapping = map[string]func(*schema.FoodEstablishmentReservationTrait, []string){}
var customFoodEstablishmentReservationTraitMapping = map[string]func(*schema.FoodEstablishmentReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FoodEstablishmentReservation", func(ctx jsonld.Context) (interface{}) {
		return NewFoodEstablishmentReservationFromContext(ctx)
	})





	customFoodEstablishmentReservationTraitMapping["http://schema.org/endTime"] = func(x *schema.FoodEstablishmentReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.EndTime = &y

		return
	}

	customFoodEstablishmentReservationTraitMapping["http://schema.org/startTime"] = func(x *schema.FoodEstablishmentReservationTrait, ctx jsonld.Context, s string){
		var y schema.DateTime
		if strings.HasPrefix(s, "_:") {
			y = NewDateTimeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDateTime()
			y.Id = s
		}

		x.StartTime = &y

		return
	}

	customFoodEstablishmentReservationTraitMapping["http://schema.org/partySize"] = func(x *schema.FoodEstablishmentReservationTrait, ctx jsonld.Context, s string){
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
}

func NewFoodEstablishmentReservationFromContext(ctx jsonld.Context) (x schema.FoodEstablishmentReservation) {
	x.Type = "http://schema.org/FoodEstablishmentReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFoodEstablishmentReservationTrait(ctx, &x.FoodEstablishmentReservationTrait)
	MapCustomToFoodEstablishmentReservationTrait(ctx, &x.FoodEstablishmentReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFoodEstablishmentReservationTrait(ctx jsonld.Context, x *schema.FoodEstablishmentReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFoodEstablishmentReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFoodEstablishmentReservationTrait(ctx jsonld.Context, x *schema.FoodEstablishmentReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFoodEstablishmentReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}