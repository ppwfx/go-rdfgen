package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBeach strings.Replacer
var strconvBeach strconv.NumError

var basicBeachTraitMapping = map[string]func(*schema.BeachTrait, []string){}
var customBeachTraitMapping = map[string]func(*schema.BeachTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Beach", func(ctx jsonld.Context) (interface{}) {
		return NewBeachFromContext(ctx)
	})

}

func NewBeachFromContext(ctx jsonld.Context) (x schema.Beach) {
	x.Type = "http://schema.org/Beach"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBeachTrait(ctx, &x.BeachTrait)
	MapCustomToBeachTrait(ctx, &x.BeachTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBeachTrait(ctx jsonld.Context, x *schema.BeachTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBeachTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBeachTrait(ctx jsonld.Context, x *schema.BeachTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBeachTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}