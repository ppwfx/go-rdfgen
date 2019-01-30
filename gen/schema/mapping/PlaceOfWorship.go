package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPlaceOfWorship strings.Replacer
var strconvPlaceOfWorship strconv.NumError

var basicPlaceOfWorshipTraitMapping = map[string]func(*schema.PlaceOfWorshipTrait, []string){}
var customPlaceOfWorshipTraitMapping = map[string]func(*schema.PlaceOfWorshipTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PlaceOfWorship", func(ctx jsonld.Context) (interface{}) {
		return NewPlaceOfWorshipFromContext(ctx)
	})

}

func NewPlaceOfWorshipFromContext(ctx jsonld.Context) (x schema.PlaceOfWorship) {
	x.Type = "http://schema.org/PlaceOfWorship"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPlaceOfWorshipTrait(ctx jsonld.Context, x *schema.PlaceOfWorshipTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPlaceOfWorshipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPlaceOfWorshipTrait(ctx jsonld.Context, x *schema.PlaceOfWorshipTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPlaceOfWorshipTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}