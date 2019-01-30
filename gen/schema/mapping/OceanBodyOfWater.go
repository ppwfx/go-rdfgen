package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOceanBodyOfWater strings.Replacer
var strconvOceanBodyOfWater strconv.NumError

var basicOceanBodyOfWaterTraitMapping = map[string]func(*schema.OceanBodyOfWaterTrait, []string){}
var customOceanBodyOfWaterTraitMapping = map[string]func(*schema.OceanBodyOfWaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OceanBodyOfWater", func(ctx jsonld.Context) (interface{}) {
		return NewOceanBodyOfWaterFromContext(ctx)
	})

}

func NewOceanBodyOfWaterFromContext(ctx jsonld.Context) (x schema.OceanBodyOfWater) {
	x.Type = "http://schema.org/OceanBodyOfWater"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOceanBodyOfWaterTrait(ctx, &x.OceanBodyOfWaterTrait)
	MapCustomToOceanBodyOfWaterTrait(ctx, &x.OceanBodyOfWaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOceanBodyOfWaterTrait(ctx jsonld.Context, x *schema.OceanBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOceanBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOceanBodyOfWaterTrait(ctx jsonld.Context, x *schema.OceanBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOceanBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}