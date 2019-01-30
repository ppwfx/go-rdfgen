package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAirport strings.Replacer
var strconvAirport strconv.NumError

var basicAirportTraitMapping = map[string]func(*schema.AirportTrait, []string){}
var customAirportTraitMapping = map[string]func(*schema.AirportTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Airport", func(ctx jsonld.Context) (interface{}) {
		return NewAirportFromContext(ctx)
	})


	basicAirportTraitMapping["http://schema.org/iataCode"] = func(x *schema.AirportTrait, s []string) {
		x.IataCode = s[0]
	}


	basicAirportTraitMapping["http://schema.org/icaoCode"] = func(x *schema.AirportTrait, s []string) {
		x.IcaoCode = s[0]
	}

}

func NewAirportFromContext(ctx jsonld.Context) (x schema.Airport) {
	x.Type = "http://schema.org/Airport"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAirportTrait(ctx, &x.AirportTrait)
	MapCustomToAirportTrait(ctx, &x.AirportTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAirportTrait(ctx jsonld.Context, x *schema.AirportTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAirportTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAirportTrait(ctx jsonld.Context, x *schema.AirportTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAirportTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}