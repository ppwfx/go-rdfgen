package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPark strings.Replacer
var strconvPark strconv.NumError

var basicParkTraitMapping = map[string]func(*schema.ParkTrait, []string){}
var customParkTraitMapping = map[string]func(*schema.ParkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Park", func(ctx jsonld.Context) (interface{}) {
		return NewParkFromContext(ctx)
	})

}

func NewParkFromContext(ctx jsonld.Context) (x schema.Park) {
	x.Type = "http://schema.org/Park"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToParkTrait(ctx, &x.ParkTrait)
	MapCustomToParkTrait(ctx, &x.ParkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToParkTrait(ctx jsonld.Context, x *schema.ParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToParkTrait(ctx jsonld.Context, x *schema.ParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}