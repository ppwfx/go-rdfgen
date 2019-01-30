package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRiverBodyOfWater strings.Replacer
var strconvRiverBodyOfWater strconv.NumError

var basicRiverBodyOfWaterTraitMapping = map[string]func(*schema.RiverBodyOfWaterTrait, []string){}
var customRiverBodyOfWaterTraitMapping = map[string]func(*schema.RiverBodyOfWaterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RiverBodyOfWater", func(ctx jsonld.Context) (interface{}) {
		return NewRiverBodyOfWaterFromContext(ctx)
	})

}

func NewRiverBodyOfWaterFromContext(ctx jsonld.Context) (x schema.RiverBodyOfWater) {
	x.Type = "http://schema.org/RiverBodyOfWater"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRiverBodyOfWaterTrait(ctx, &x.RiverBodyOfWaterTrait)
	MapCustomToRiverBodyOfWaterTrait(ctx, &x.RiverBodyOfWaterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRiverBodyOfWaterTrait(ctx jsonld.Context, x *schema.RiverBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRiverBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRiverBodyOfWaterTrait(ctx jsonld.Context, x *schema.RiverBodyOfWaterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRiverBodyOfWaterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}