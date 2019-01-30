package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusStation strings.Replacer
var strconvBusStation strconv.NumError

var basicBusStationTraitMapping = map[string]func(*schema.BusStationTrait, []string){}
var customBusStationTraitMapping = map[string]func(*schema.BusStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusStation", func(ctx jsonld.Context) (interface{}) {
		return NewBusStationFromContext(ctx)
	})

}

func NewBusStationFromContext(ctx jsonld.Context) (x schema.BusStation) {
	x.Type = "http://schema.org/BusStation"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusStationTrait(ctx, &x.BusStationTrait)
	MapCustomToBusStationTrait(ctx, &x.BusStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusStationTrait(ctx jsonld.Context, x *schema.BusStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusStationTrait(ctx jsonld.Context, x *schema.BusStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}