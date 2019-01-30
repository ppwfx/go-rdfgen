package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMosque strings.Replacer
var strconvMosque strconv.NumError

var basicMosqueTraitMapping = map[string]func(*schema.MosqueTrait, []string){}
var customMosqueTraitMapping = map[string]func(*schema.MosqueTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Mosque", func(ctx jsonld.Context) (interface{}) {
		return NewMosqueFromContext(ctx)
	})

}

func NewMosqueFromContext(ctx jsonld.Context) (x schema.Mosque) {
	x.Type = "http://schema.org/Mosque"
	MapBasicToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)
	MapCustomToPlaceOfWorshipTrait(ctx, &x.PlaceOfWorshipTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMosqueTrait(ctx, &x.MosqueTrait)
	MapCustomToMosqueTrait(ctx, &x.MosqueTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMosqueTrait(ctx jsonld.Context, x *schema.MosqueTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMosqueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMosqueTrait(ctx jsonld.Context, x *schema.MosqueTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMosqueTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}