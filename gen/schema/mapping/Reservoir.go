package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReservoir strings.Replacer
var strconvReservoir strconv.NumError

var basicReservoirTraitMapping = map[string]func(*schema.ReservoirTrait, []string){}
var customReservoirTraitMapping = map[string]func(*schema.ReservoirTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Reservoir", func(ctx jsonld.Context) (interface{}) {
		return NewReservoirFromContext(ctx)
	})

}

func NewReservoirFromContext(ctx jsonld.Context) (x schema.Reservoir) {
	x.Type = "http://schema.org/Reservoir"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReservoirTrait(ctx, &x.ReservoirTrait)
	MapCustomToReservoirTrait(ctx, &x.ReservoirTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReservoirTrait(ctx jsonld.Context, x *schema.ReservoirTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReservoirTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReservoirTrait(ctx jsonld.Context, x *schema.ReservoirTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReservoirTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}