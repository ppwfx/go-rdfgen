package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWaterfall strings.Replacer
var strconvWaterfall strconv.NumError

var basicWaterfallTraitMapping = map[string]func(*schema.WaterfallTrait, []string){}
var customWaterfallTraitMapping = map[string]func(*schema.WaterfallTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Waterfall", func(ctx jsonld.Context) (interface{}) {
		return NewWaterfallFromContext(ctx)
	})

}

func NewWaterfallFromContext(ctx jsonld.Context) (x schema.Waterfall) {
	x.Type = "http://schema.org/Waterfall"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWaterfallTrait(ctx, &x.WaterfallTrait)
	MapCustomToWaterfallTrait(ctx, &x.WaterfallTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWaterfallTrait(ctx jsonld.Context, x *schema.WaterfallTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWaterfallTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWaterfallTrait(ctx jsonld.Context, x *schema.WaterfallTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWaterfallTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}