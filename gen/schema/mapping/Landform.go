package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLandform strings.Replacer
var strconvLandform strconv.NumError

var basicLandformTraitMapping = map[string]func(*schema.LandformTrait, []string){}
var customLandformTraitMapping = map[string]func(*schema.LandformTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Landform", func(ctx jsonld.Context) (interface{}) {
		return NewLandformFromContext(ctx)
	})

}

func NewLandformFromContext(ctx jsonld.Context) (x schema.Landform) {
	x.Type = "http://schema.org/Landform"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLandformTrait(ctx jsonld.Context, x *schema.LandformTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLandformTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLandformTrait(ctx jsonld.Context, x *schema.LandformTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLandformTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}