package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEventReservation strings.Replacer
var strconvEventReservation strconv.NumError

var basicEventReservationTraitMapping = map[string]func(*schema.EventReservationTrait, []string){}
var customEventReservationTraitMapping = map[string]func(*schema.EventReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/EventReservation", func(ctx jsonld.Context) (interface{}) {
		return NewEventReservationFromContext(ctx)
	})

}

func NewEventReservationFromContext(ctx jsonld.Context) (x schema.EventReservation) {
	x.Type = "http://schema.org/EventReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEventReservationTrait(ctx, &x.EventReservationTrait)
	MapCustomToEventReservationTrait(ctx, &x.EventReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEventReservationTrait(ctx jsonld.Context, x *schema.EventReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEventReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEventReservationTrait(ctx jsonld.Context, x *schema.EventReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEventReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}