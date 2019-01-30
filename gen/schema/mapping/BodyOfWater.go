package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBodyOfWater strings.Replacer
var strconvBodyOfWater strconv.NumError

var basicBodyOfWaterTraitMapping = map[string]func(*schema.BodyOfWaterTrait, []string){}
var customBodyOfWaterTraitMapping = map[string]func(*schema.BodyOfWaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BodyOfWater", func(ctx jsonld.Context) (interface{}) {
		return NewBodyOfWaterFromContext(ctx)
	})

}

func NewBodyOfWaterFromContext(ctx jsonld.Context) (x schema.BodyOfWater) {
	x.Type = "http://schema.org/BodyOfWater"
	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBodyOfWaterTrait(ctx jsonld.Context, x *schema.BodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBodyOfWaterTrait(ctx jsonld.Context, x *schema.BodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}