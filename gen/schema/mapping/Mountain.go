package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMountain strings.Replacer
var strconvMountain strconv.NumError

var basicMountainTraitMapping = map[string]func(*schema.MountainTrait, []string){}
var customMountainTraitMapping = map[string]func(*schema.MountainTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Mountain", func(ctx jsonld.Context) (interface{}) {
		return NewMountainFromContext(ctx)
	})

}

func NewMountainFromContext(ctx jsonld.Context) (x schema.Mountain) {
	x.Type = "http://schema.org/Mountain"
	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMountainTrait(ctx, &x.MountainTrait)
	MapCustomToMountainTrait(ctx, &x.MountainTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMountainTrait(ctx jsonld.Context, x *schema.MountainTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMountainTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMountainTrait(ctx jsonld.Context, x *schema.MountainTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMountainTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}