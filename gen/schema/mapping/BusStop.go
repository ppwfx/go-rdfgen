package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusStop strings.Replacer
var strconvBusStop strconv.NumError

var basicBusStopTraitMapping = map[string]func(*schema.BusStopTrait, []string){}
var customBusStopTraitMapping = map[string]func(*schema.BusStopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusStop", func(ctx jsonld.Context) (interface{}) {
		return NewBusStopFromContext(ctx)
	})

}

func NewBusStopFromContext(ctx jsonld.Context) (x schema.BusStop) {
	x.Type = "http://schema.org/BusStop"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusStopTrait(ctx, &x.BusStopTrait)
	MapCustomToBusStopTrait(ctx, &x.BusStopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusStopTrait(ctx jsonld.Context, x *schema.BusStopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusStopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusStopTrait(ctx jsonld.Context, x *schema.BusStopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusStopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}