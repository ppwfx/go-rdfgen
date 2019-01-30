package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTrainReservation strings.Replacer
var strconvTrainReservation strconv.NumError

var basicTrainReservationTraitMapping = map[string]func(*schema.TrainReservationTrait, []string){}
var customTrainReservationTraitMapping = map[string]func(*schema.TrainReservationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TrainReservation", func(ctx jsonld.Context) (interface{}) {
		return NewTrainReservationFromContext(ctx)
	})

}

func NewTrainReservationFromContext(ctx jsonld.Context) (x schema.TrainReservation) {
	x.Type = "http://schema.org/TrainReservation"
	MapBasicToReservationTrait(ctx, &x.ReservationTrait)
	MapCustomToReservationTrait(ctx, &x.ReservationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTrainReservationTrait(ctx, &x.TrainReservationTrait)
	MapCustomToTrainReservationTrait(ctx, &x.TrainReservationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTrainReservationTrait(ctx jsonld.Context, x *schema.TrainReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTrainReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTrainReservationTrait(ctx jsonld.Context, x *schema.TrainReservationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTrainReservationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}