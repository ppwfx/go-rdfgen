package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCanal strings.Replacer
var strconvCanal strconv.NumError

var basicCanalTraitMapping = map[string]func(*schema.CanalTrait, []string){}
var customCanalTraitMapping = map[string]func(*schema.CanalTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Canal", func(ctx jsonld.Context) (interface{}) {
		return NewCanalFromContext(ctx)
	})

}

func NewCanalFromContext(ctx jsonld.Context) (x schema.Canal) {
	x.Type = "http://schema.org/Canal"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCanalTrait(ctx, &x.CanalTrait)
	MapCustomToCanalTrait(ctx, &x.CanalTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCanalTrait(ctx jsonld.Context, x *schema.CanalTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCanalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCanalTrait(ctx jsonld.Context, x *schema.CanalTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCanalTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}