package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusReservation strings.Replacer
var strconvBusReservation strconv.NumError

var basicBusReservationTraitMapping = map[string]func(*schema.BusReservationTrait, []string){}
var customBusReservationTraitMapping = map[string]func(*schema.BusReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusReservation", func(ctx jsonld.Context) (interface{}) {
		return NewBusReservationFromContext(ctx)
	})

}

func NewBusReservationFromContext(ctx jsonld.Context) (x schema.BusReservation) {
	x.Type = "http://schema.org/BusReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusReservationTrait(ctx, &x.BusReservationTrait)
	MapCustomToBusReservationTrait(ctx, &x.BusReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusReservationTrait(ctx jsonld.Context, x *schema.BusReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusReservationTrait(ctx jsonld.Context, x *schema.BusReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}