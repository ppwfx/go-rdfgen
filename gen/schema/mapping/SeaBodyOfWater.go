package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSeaBodyOfWater strings.Replacer
var strconvSeaBodyOfWater strconv.NumError

var basicSeaBodyOfWaterTraitMapping = map[string]func(*schema.SeaBodyOfWaterTrait, []string){}
var customSeaBodyOfWaterTraitMapping = map[string]func(*schema.SeaBodyOfWaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SeaBodyOfWater", func(ctx jsonld.Context) (interface{}) {
		return NewSeaBodyOfWaterFromContext(ctx)
	})

}

func NewSeaBodyOfWaterFromContext(ctx jsonld.Context) (x schema.SeaBodyOfWater) {
	x.Type = "http://schema.org/SeaBodyOfWater"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSeaBodyOfWaterTrait(ctx, &x.SeaBodyOfWaterTrait)
	MapCustomToSeaBodyOfWaterTrait(ctx, &x.SeaBodyOfWaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSeaBodyOfWaterTrait(ctx jsonld.Context, x *schema.SeaBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSeaBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSeaBodyOfWaterTrait(ctx jsonld.Context, x *schema.SeaBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSeaBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}