package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLakeBodyOfWater strings.Replacer
var strconvLakeBodyOfWater strconv.NumError

var basicLakeBodyOfWaterTraitMapping = map[string]func(*schema.LakeBodyOfWaterTrait, []string){}
var customLakeBodyOfWaterTraitMapping = map[string]func(*schema.LakeBodyOfWaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LakeBodyOfWater", func(ctx jsonld.Context) (interface{}) {
		return NewLakeBodyOfWaterFromContext(ctx)
	})

}

func NewLakeBodyOfWaterFromContext(ctx jsonld.Context) (x schema.LakeBodyOfWater) {
	x.Type = "http://schema.org/LakeBodyOfWater"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLakeBodyOfWaterTrait(ctx, &x.LakeBodyOfWaterTrait)
	MapCustomToLakeBodyOfWaterTrait(ctx, &x.LakeBodyOfWaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLakeBodyOfWaterTrait(ctx jsonld.Context, x *schema.LakeBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLakeBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLakeBodyOfWaterTrait(ctx jsonld.Context, x *schema.LakeBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLakeBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}