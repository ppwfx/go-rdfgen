package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCrematorium strings.Replacer
var strconvCrematorium strconv.NumError

var basicCrematoriumTraitMapping = map[string]func(*schema.CrematoriumTrait, []string){}
var customCrematoriumTraitMapping = map[string]func(*schema.CrematoriumTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Crematorium", func(ctx jsonld.Context) (interface{}) {
		return NewCrematoriumFromContext(ctx)
	})

}

func NewCrematoriumFromContext(ctx jsonld.Context) (x schema.Crematorium) {
	x.Type = "http://schema.org/Crematorium"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCrematoriumTrait(ctx, &x.CrematoriumTrait)
	MapCustomToCrematoriumTrait(ctx, &x.CrematoriumTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCrematoriumTrait(ctx jsonld.Context, x *schema.CrematoriumTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCrematoriumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCrematoriumTrait(ctx jsonld.Context, x *schema.CrematoriumTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCrematoriumTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}