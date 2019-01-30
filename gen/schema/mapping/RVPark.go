package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRVPark strings.Replacer
var strconvRVPark strconv.NumError

var basicRVParkTraitMapping = map[string]func(*schema.RVParkTrait, []string){}
var customRVParkTraitMapping = map[string]func(*schema.RVParkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RVPark", func(ctx jsonld.Context) (interface{}) {
		return NewRVParkFromContext(ctx)
	})

}

func NewRVParkFromContext(ctx jsonld.Context) (x schema.RVPark) {
	x.Type = "http://schema.org/RVPark"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRVParkTrait(ctx, &x.RVParkTrait)
	MapCustomToRVParkTrait(ctx, &x.RVParkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRVParkTrait(ctx jsonld.Context, x *schema.RVParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRVParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRVParkTrait(ctx jsonld.Context, x *schema.RVParkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRVParkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}