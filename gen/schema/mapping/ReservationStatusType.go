package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReservationStatusType strings.Replacer
var strconvReservationStatusType strconv.NumError

var basicReservationStatusTypeTraitMapping = map[string]func(*schema.ReservationStatusTypeTrait, []string){}
var customReservationStatusTypeTraitMapping = map[string]func(*schema.ReservationStatusTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReservationStatusType", func(ctx jsonld.Context) (interface{}) {
		return NewReservationStatusTypeFromContext(ctx)
	})

}

func NewReservationStatusTypeFromContext(ctx jsonld.Context) (x schema.ReservationStatusType) {
	x.Type = "http://schema.org/ReservationStatusType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReservationStatusTypeTrait(ctx, &x.ReservationStatusTypeTrait)
	MapCustomToReservationStatusTypeTrait(ctx, &x.ReservationStatusTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReservationStatusTypeTrait(ctx jsonld.Context, x *schema.ReservationStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReservationStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReservationStatusTypeTrait(ctx jsonld.Context, x *schema.ReservationStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReservationStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}