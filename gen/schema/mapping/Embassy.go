package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEmbassy strings.Replacer
var strconvEmbassy strconv.NumError

var basicEmbassyTraitMapping = map[string]func(*schema.EmbassyTrait, []string){}
var customEmbassyTraitMapping = map[string]func(*schema.EmbassyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Embassy", func(ctx jsonld.Context) (interface{}) {
		return NewEmbassyFromContext(ctx)
	})

}

func NewEmbassyFromContext(ctx jsonld.Context) (x schema.Embassy) {
	x.Type = "http://schema.org/Embassy"
	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEmbassyTrait(ctx, &x.EmbassyTrait)
	MapCustomToEmbassyTrait(ctx, &x.EmbassyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEmbassyTrait(ctx jsonld.Context, x *schema.EmbassyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEmbassyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEmbassyTrait(ctx jsonld.Context, x *schema.EmbassyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEmbassyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}