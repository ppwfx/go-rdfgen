package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCemetery strings.Replacer
var strconvCemetery strconv.NumError

var basicCemeteryTraitMapping = map[string]func(*schema.CemeteryTrait, []string){}
var customCemeteryTraitMapping = map[string]func(*schema.CemeteryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Cemetery", func(ctx jsonld.Context) (interface{}) {
		return NewCemeteryFromContext(ctx)
	})

}

func NewCemeteryFromContext(ctx jsonld.Context) (x schema.Cemetery) {
	x.Type = "http://schema.org/Cemetery"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCemeteryTrait(ctx, &x.CemeteryTrait)
	MapCustomToCemeteryTrait(ctx, &x.CemeteryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCemeteryTrait(ctx jsonld.Context, x *schema.CemeteryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCemeteryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCemeteryTrait(ctx jsonld.Context, x *schema.CemeteryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCemeteryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}